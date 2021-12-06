package puzzle3

import (
	"strings"
	"strconv"
	"fmt"
	"github.com/crabbey/aoc2021/common"
	"github.com/davecgh/go-spew/spew"
)

var _ = spew.Dump
var _ = fmt.Println
var _ = strings.Split

type Puzzle3 struct {
	Lines map[int]Line
}

type Line map[int]int

func (l Line) String() string {
	ret := ""
	for x := 0; x < len(l); x++ {
		ret += strconv.Itoa(l[x])
	}
	return ret
}

func (p *Puzzle3) Parse(input []string) {
	p.Lines = make(map[int]Line)
	for x, l := range input {
		line := make(map[int]int)
		parts := strings.Split(l, "")
		for y, c := range parts {
			line[y], _ = strconv.Atoi(c)
		}
		p.Lines[x] = line
	}
}

func (p *Puzzle3) Interpret1() int64 {
	gamma := ""
	epsilon := ""
	for y := 0; y < len(p.Lines[0]); y++ {  // y = column
		count := 0
		for x := 0; x < len(p.Lines); x++ { // x = row
			if p.Lines[x][y] == 0 {
				count -= 1
			} else {
				count += 1
			}
		}
		if count > 0 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	g, _ := strconv.ParseInt(gamma, 2, 0)
	e, _ := strconv.ParseInt(epsilon, 2, 0)
	fmt.Printf("%v, %v, %v, %v\n", gamma, epsilon, g, e)
	return g * e
}

func (p *Puzzle3) Interpret2() int64 {
	oxy := "0"
	co2 := "0"
	lines := make(map[int]Line)
	for k, v := range p.Lines {
	  lines[k] = v
	}
	bitwidth := len(lines[0])
	for y := 0; y < bitwidth; y++ {  // y = column
		count := 0
		for x, _ := range lines { // x = row
			if lines[x][y] == 0 {
				count -= 1
			} else {
				count += 1
			}
		}
		for x, _ := range lines {
			if count >= 0 {
				if lines[x][y] == 0 {
					delete(lines, x)
				}
			} else {
				if lines[x][y] == 1 {
					delete(lines, x)
				}
			}
		}
		// fmt.Printf("%v lines remaining\n", len(lines))
		// for _, l := range lines {
		// 	fmt.Printf("%v, ", l.String())
		// }
		// fmt.Println()
		if len(lines) == 1 {
			break;
		}
	}
	for _, l := range lines {
		oxy = l.String()
	}

	for k, v := range p.Lines {
	  lines[k] = v
	}
	for y := 0; y < bitwidth; y++ {  // y = column
		count := 0
		for x, _ := range lines { // x = row
			if lines[x][y] == 0 {
				count -= 1
			} else {
				count += 1
			}
		}
		for x, _ := range lines {
			if count >= 0 {
				if lines[x][y] == 1 {
					delete(lines, x)
				}
			} else {
				if lines[x][y] == 0 {
					delete(lines, x)
				}
			}
		}
		// fmt.Printf("%v lines remaining\n", len(lines))
		// for _, l := range lines {
			// fmt.Printf("%v, ", l.String())
		// }
		// fmt.Println()
		if len(lines) == 1 {
			break;
		}
	}
	for _, l := range lines {
		co2 = l.String()
	}
	o, _ := strconv.ParseInt(oxy, 2, 0)
	c, _ := strconv.ParseInt(co2, 2, 0)
	fmt.Printf("%v, %v, %v, %v\n", oxy, co2, o, c)
	return o * c
}

func (p Puzzle3) Part1(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		return nil, err
	}
	output := common.NewSolution(input, "")
	p.Parse(i)
	answer := p.Interpret1()
	output.Text = strconv.FormatInt(answer, 10)
	return output, nil
}

func (p Puzzle3) Part2(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		return nil, err
	}
	output := common.NewSolution(input, "")
	p.Parse(i)
	answer := p.Interpret2()
	output.Text = strconv.FormatInt(answer, 10)
	return output, nil
}