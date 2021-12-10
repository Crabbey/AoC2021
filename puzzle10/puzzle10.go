package puzzle10

import (
	"strings"
	"strconv"
	"fmt"
	"regexp"
	"sort"
	"github.com/crabbey/aoc2021/common"
	"github.com/davecgh/go-spew/spew"
)

var _ = spew.Dump
var _ = strings.Split
var _ = strconv.Itoa

type Puzzle10 struct {

}

var ChunkPairs = map[string]string{
	"{": "}",
	"[": "]",
	"<": ">",
	"(": ")",
}

var errorVals = map[string]int{
	"}": 1197,
	"]": 57,
	">": 25137,
	")": 3,
}

var autocompleteVals = map[string]int{
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
}

func (p Puzzle10) Part1(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return nil, err
	}
	output := common.NewSolution(input, "")
	readyChunks, _ := regexp.Compile(`(\{\}|\(\)|\[\]|\<\>)`)
	invalid, _ := regexp.Compile(`(\{\>|\{\]|\{\)|\[\}|\[\>|\[\)|\(\]|\(\>|\(\}|\<\}|\<\]|\<\))`)
	total := 0
	for _, l := range i {
		newL := ""
		for {
			newL = readyChunks.ReplaceAllString(l, "")
			if newL == l {
				// No change
				break
			}
			l = newL
		}
		invalidChunks := invalid.FindAllString(newL, -1)
		if len(invalidChunks) > 0 {
			for _, chunk := range invalidChunks {
				c := strings.Split(chunk, "")
				total += errorVals[c[1]]
			}
		}
	}

	output.Text = fmt.Sprintf("%v", total)
	return output, nil
}

func (p Puzzle10) Part2(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return nil, err
	}
	output := common.NewSolution(input, "")
	readyChunks, _ := regexp.Compile(`(\{\}|\(\)|\[\]|\<\>)`)
	invalid, _ := regexp.Compile(`(\{\>|\{\]|\{\)|\[\}|\[\>|\[\)|\(\]|\(\>|\(\}|\<\}|\<\]|\<\))`)
	var scores []int
	for _, l := range i {
		total := 0
		newL := ""
		for {
			newL = readyChunks.ReplaceAllString(l, "")
			if newL == l {
				// No change
				break
			}
			l = newL
		}
		invalidChunks := invalid.FindAllString(newL, -1)
		if len(invalidChunks) == 0 {
			// Valid, but incomplete
			for {
				lastChar := l[len(l)-1:]
				total *= 5 
				total += autocompleteVals[ChunkPairs[lastChar]]
				// fmt.Print(ChunkPairs[lastChar])
				l = l[:len(l)-1]
				if l == "" {
					break
				}
			}
			// fmt.Print(" - ")
			// fmt.Println(total)
			scores = append(scores, total)
		}
	}
	sort.Ints(scores)
	middle := scores[(len(scores)-1)/2]
	output.Text = fmt.Sprintf("%v", middle)
	return output, nil
}