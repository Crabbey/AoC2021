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
	Round int
}

type Instruction struct {
	Pair string
	Value string
	Re *regexp.Regexp
}

func (p *Puzzle14) AddInstruction(i string) {
	parts := strings.Split(i, " -> ")
	inst := &Instruction{
		Pair: parts[0],
		Value: parts[1],
		Re: regexp.MustCompile(parts[0]),
	}
	p.Instructions = append(p.Instructions, inst)
}

func (p *Puzzle14) Parse(i []string) {
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
}
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

type Match struct {
	X int
	Val string
}

func (p *Puzzle14) DoRound() {
	p.Round++
	additions := make(map[int]string)
	for a, i := range p.Instructions {
		fmt.Printf("R %v I %v\n", p.Round, a)
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
	// spew.Dump(p.State)
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
		p.DoRound()
	}
	output.Text = fmt.Sprintf("%v", p.CountCommonality())
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
		p.DoRound()
	}
	output.Text = fmt.Sprintf("%v", p.CountCommonality())
	return output, nil
}