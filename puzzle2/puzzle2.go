package puzzle2

import (
	"strconv"
	"strings"
	"fmt"
	// "regexp"
	"github.com/crabbey/aoc2021/common"
	"github.com/davecgh/go-spew/spew"
)

var _ = spew.Dump
var _ = fmt.Println

type Puzzle2 struct {
	entries map[int]*Entry
	Pos struct {
		Horizontal int
		Depth int
		Aim int
	}
}

type Entry struct {
	Command string
	Value int
}


func (p *Puzzle2) Parse(input []string) {
	p.entries = make(map[int]*Entry)
	for x, l := range input {
		parts := strings.Split(l, " ")
		val, _ := strconv.Atoi(parts[1])
		p.entries[x] = &Entry{
			Command: parts[0],
			Value: val,
		}
	}
}


func (p *Puzzle2) Execute1(e *Entry) {
	switch e.Command {
	case "up":
		p.Pos.Depth -= e.Value
	case "down":
		p.Pos.Depth += e.Value
	case "forward":
		p.Pos.Horizontal += e.Value
	}
}


func (p *Puzzle2) Execute2(e *Entry) {
	switch e.Command {
	case "up":
		p.Pos.Aim -= e.Value
	case "down":
		p.Pos.Aim += e.Value
	case "forward":
		p.Pos.Depth += p.Pos.Aim * e.Value
		p.Pos.Horizontal += e.Value
	}
	// fmt.Printf("'%v %v' => Horizontal: %v, Depth: %v, Aim: %v\n", e.Command, e.Value, p.Pos.Horizontal, p.Pos.Depth, p.Pos.Aim)
}

func (p Puzzle2) Part1(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		return nil, err
	}
	output := common.NewSolution(input, "")
	p.Parse(i)
	for i := 0; i < len(p.entries); i++ {
		p.Execute1(p.entries[i])
	}
	output.Text = strconv.Itoa(p.Pos.Horizontal * p.Pos.Depth)
	return output, nil
}

func (p Puzzle2) Part2(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		return nil, err
	}
	output := common.NewSolution(input, "")
	p.Parse(i)
	for i := 0; i < len(p.entries); i++ {
		p.Execute2(p.entries[i])
	}
	output.Text = strconv.Itoa(p.Pos.Horizontal * p.Pos.Depth)
	return output, nil
}