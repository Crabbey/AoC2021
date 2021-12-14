package puzzle14

import (
	"strings"
	"strconv"
	"fmt"
	"sort"
	"regexp"
	"sync"
	"github.com/crabbey/aoc2021/common"
	"github.com/davecgh/go-spew/spew"
)

var _ = spew.Dump
var _ = strings.Split
var _ = strconv.Itoa

type Puzzle14 struct {
	State string
	Instructions []*Instruction
	Table map[string]int
	Round int
}

type Instruction struct {
	Pair string
	Value string
	Re *regexp.Regexp
	Products []string
}

func (p *Puzzle14) AddInstruction(i string) {
	parts := strings.Split(i, " -> ")
	inst := &Instruction{
		Pair: parts[0],
		Value: parts[1],
		Re: regexp.MustCompile(parts[0]),
		Products: []string{
			string(parts[0][0]) + parts[1],
			parts[1] + string(parts[0][1]),
		},
	}
	for _, v := range inst.Products {
		p.Table[v] = 0
	}
	p.Table[inst.Pair] = 0
	p.Instructions = append(p.Instructions, inst)
}

func (p *Puzzle14) Parse(i []string) {
	p.Table = make(map[string]int)
	isInstructions := false
	for _, l := range i {
		if l == "" {
			isInstructions = true
			continue
		}
		if isInstructions {
			p.AddInstruction(l)
			continue
		}
		p.State = l
	}
	for x := 0; x < len(p.State)-1; x++ {
		p.Table[p.State[x:x+2]]++
	}

	// test := []string{
	// 	"NN",
	// 	"NC",
	// 	"CB",
	// }
	// for _, v := range test {
	// 	fmt.Printf("%v: %v\n", v, p.Table[v])
	// }
	// fmt.Println()
}

func (p *Puzzle14) DoRoundTable() {
	p.Round++
	newTable := make(map[string]int)
	for _, i := range p.Instructions {
		for _, x := range i.Products {
			newTable[x] += p.Table[i.Pair]
		}
	}
	p.Table = newTable
	// test := []string{
	// 	"NC",
	// 	"CN",
	// 	"NB",
	// 	"BC",
	// 	"CH",
	// 	"HB",
	// }
	// for _, v := range test {
	// 	fmt.Printf("%v: %v\n", v, p.Table[v])
	// }
}

func (p *Puzzle14) CountCommonalityTable() int {
	vals := make(map[string]int)
	for k, v := range p.Table {
		vals[string(k[0])] += v
		vals[string(k[1])] += v
	}

	// The first and last letter are off-by-one
	vals[string(p.State[0])]++
	vals[string(p.State[len(p.State)-1])]++

	// All values are doubled, because they're accounted for twice
	for k, v := range vals {
		vals[k] = v/2
	}

	// spew.Dump(vals)
	highest := 0
	lowest := 100000000000000000 // big for init
	for _	, v := range vals {
		if v > highest {
			highest = v
		} 
		if v < lowest {
			lowest = v
		}
	}
	return highest - lowest
}

// Deprecated
func (p *Puzzle14) CountCommonality() int {
	vals := make(map[string]int)
	for _, v := range p.State {
		if _, ok := vals[string(v)]; !ok {
			vals[string(v)] = 0;
		}
		vals[string(v)]++
	}
	highest := 0
	lowest := 100000000000000000 // big for init
	for _	, v := range vals {
		if v > highest {
			highest = v
		} 
		if v < lowest {
			lowest = v
		}
	}
	return highest - lowest
}

// Deprecated
type Match struct {
	X int
	Val string
}

// Deprecated
func (p *Puzzle14) DoRound() {
	p.Round++
	additions := make(map[int]string)
	for _, i := range p.Instructions {
		wg := &sync.WaitGroup{}
		ch := make(chan *Match)
		wg.Add(len(p.State)-1)
		go func(resp chan *Match, wg *sync.WaitGroup) {
			wg.Wait()
			close(resp)
		}(ch, wg)
		for x := 0; x < len(p.State)-1; x++ {
			go func(resp chan *Match, wg *sync.WaitGroup, x int, i *Instruction) {
				res := strings.Index(p.State[x:x+2], i.Pair)
				if res == -1 {
					wg.Done()
					return;
				}
				resp <- &Match{
					X: x,
					Val: i.Value,
				}
				wg.Done()
			}(ch, wg, x, i)
		}
		for res := range ch {
			additions[res.X] = res.Val
		}
	}
	keys := []int{}
	for a, _ := range additions {
		keys = append(keys, a)
	}
	sort.Ints(keys)
	for a, k := range keys {
		p.State = p.State[:k+a] + additions[k] + p.State[k+a:]
	}
}

func (p Puzzle14) Part1(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return nil, err
	}
	output := common.NewSolution(input, "")
	p.Parse(i)
	for x := 0; x < 10; x++ {
		p.DoRoundTable()
	}
	output.Text = fmt.Sprintf("%v", p.CountCommonalityTable())
	return output, nil
}

func (p Puzzle14) Part2(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return nil, err
	}
	output := common.NewSolution(input, "")
	p.Parse(i)
	for x := 0; x < 40; x++ {
		p.DoRoundTable()
	}
	output.Text = fmt.Sprintf("%v", p.CountCommonalityTable())
	return output, nil
}