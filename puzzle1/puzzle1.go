package puzzle1

import (
	"github.com/crabbey/aoc2021/common"
	"github.com/davecgh/go-spew/spew"
	"strconv"
	"fmt"
)

var _, _ = spew.Dump, fmt.Println

type Puzzle1 struct {

}

func (p Puzzle1) Part1(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		return nil, err
	}
	output := common.NewSolution(input, "")
	last := int64(-1)
	incCount := 0
	for _, input1 := range i {
		val, _ := strconv.ParseInt(input1, 10, 0)
		if last == -1 {
			last = val
			continue
		}
		if last < val {
			incCount++
		}
		last = val
	}
	output.Text = strconv.Itoa(incCount)
	return output, nil
}

func (p Puzzle1) Part2(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		return nil, err
	}
	output := common.NewSolution(input, "")
	entries := make(map[int]int64)
	for x, input1 := range i {
		val, _ := strconv.ParseInt(input1, 10, 0)
		entries[x] = val
	}
	incCount := 0
	for x := 1; x < len(entries) - 2; x++ {
		this := entries[x] + entries[x+1] + entries[x+2]
		previous := entries[x-1] + entries[x] + entries[x+1]
		if this > previous {
			incCount++
		}
	}
	output.Text = strconv.Itoa(incCount)
	return output, nil
}