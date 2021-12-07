package puzzle7

import (
	"strings"
	"strconv"
	"fmt"
	"github.com/crabbey/aoc2021/common"
	"github.com/davecgh/go-spew/spew"
)

var _ = spew.Dump
var _ = strings.Split
var _ = strconv.Itoa

type Puzzle7 struct {
	Crabs []*Crab
	Highest int
	BestFuel int
	BestPosition int
}

type Crab struct {
	Position int
}

func DetermineDistance(c *Crab, endPos int) int {
	return Abs(endPos - c.Position)
}

func DetermineExpoDistance(c *Crab, endPos int) int {
	linear := DetermineDistance(c, endPos)
	return linear * (linear + 1) / 2
}

func (p *Puzzle7) FindBest(f func(*Crab, int)(int)) {
	for x := 0; x <= p.Highest; x++ {
		runFuel := 0
		for _, c := range p.Crabs {
			runFuel += f(c, x)
		}
		if runFuel < p.BestFuel || p.BestFuel == 0 {
			p.BestFuel = runFuel
			p.BestPosition = x
		}
	}
}

func (p *Puzzle7) LoadCrabbos(i []string) {
	crabs := strings.Split(i[0], ",")
	for _, x := range crabs {
		d, _ := strconv.Atoi(x)
		crab := Crab{
			Position: d,
		}
		p.Crabs = append(p.Crabs, &crab)
		if d > p.Highest {
			p.Highest = d
		}
	}
}

func (p Puzzle7) Part1(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return nil, err
	}
	output := common.NewSolution(input, "")
	p.LoadCrabbos(i)
	p.FindBest(DetermineDistance)
	output.Text = fmt.Sprintf("%v", p.BestFuel)
	return output, nil
}

func (p Puzzle7) Part2(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return nil, err
	}
	output := common.NewSolution(input, "")
	p.LoadCrabbos(i)
	p.FindBest(DetermineExpoDistance)
	output.Text = fmt.Sprintf("%v", p.BestFuel)
	return output, nil
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}