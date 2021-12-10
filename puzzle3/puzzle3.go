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
	Grid *common.IntGrid
}

func RowToString(r *common.IntRow) string {
	ret := ""
	for x := 0; x < len(r.Cols); x++ {
		ret += strconv.Itoa(r.Cols[x])
	}
	return ret
}

func (p *Puzzle3) Interpret1() int64 {
	gamma := ""
	epsilon := ""
	for y := 0; y < len(p.Grid.Rows[0].Cols); y++ {  // y = column
		count := 0
		for x := 0; x < len(p.Grid.Rows); x++ { // x = row
			if p.Grid.Rows[x].Cols[y] == 0 {
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
	// fmt.Printf("%v, %v, %v, %v\n", gamma, epsilon, g, e)
	return g * e
}

func (p *Puzzle3) Interpret2(input []string) int64 {
	oxy := "0"
	co2 := "0"

	p.Grid = common.NewIntGrid()
	p.Grid.LoadFromFile(input, func(c string)int{
		height, _ := strconv.Atoi(c)
		return height
	})

	_, bitwidth := p.Grid.Dimensions()

	for y := 0; y < bitwidth; y++ {  // y = column
		count := 0
		for x, _ := range p.Grid.Rows { // x = row
			if p.Grid.Rows[x].Cols[y] == 0 {
				count -= 1
			} else {
				count += 1
			}
		}
		for x, _ := range p.Grid.Rows {
			if count >= 0 {
				if p.Grid.Rows[x].Cols[y] == 0 {
					delete(p.Grid.Rows, x)
				}
			} else {
				if p.Grid.Rows[x].Cols[y] == 1 {
					delete(p.Grid.Rows, x)
				}
			}
		}
		if len(p.Grid.Rows) == 1 {
			break;
		}
	}
	for _, r := range p.Grid.Rows {
		oxy = RowToString(r)
	}

	p.Grid.LoadFromFile(input, func(c string)int{
		height, _ := strconv.Atoi(c)
		return height
	})
	for y := 0; y < bitwidth; y++ {  // y = column
		count := 0
		for x, _ := range p.Grid.Rows { // x = row
			if p.Grid.Rows[x].Cols[y] == 0 {
				count -= 1
			} else {
				count += 1
			}
		}
		for x, _ := range p.Grid.Rows {
			if count >= 0 {
				if p.Grid.Rows[x].Cols[y] == 1 {
					delete(p.Grid.Rows, x)
				}
			} else {
				if p.Grid.Rows[x].Cols[y] == 0 {
					delete(p.Grid.Rows, x)
				}
			}
		}
		if len(p.Grid.Rows) == 1 {
			break;
		}
	}
	for _, l := range p.Grid.Rows {
		co2 = RowToString(l)
	}
	o, _ := strconv.ParseInt(oxy, 2, 0)
	c, _ := strconv.ParseInt(co2, 2, 0)
	// fmt.Printf("%v, %v, %v, %v\n", oxy, co2, o, c)
	return o * c
}

func (p Puzzle3) Part1(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		return nil, err
	}
	output := common.NewSolution(input, "")
	p.Grid = common.NewIntGrid()
	p.Grid.LoadFromFile(i, func(c string)int{
		height, _ := strconv.Atoi(c)
		return height
	})
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
	answer := p.Interpret2(i)
	output.Text = strconv.FormatInt(answer, 10)
	return output, nil
}