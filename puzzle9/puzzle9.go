package puzzle9

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

var dirs = []string{
	"right",
	"left",
	"up",
	"down",
}


type Puzzle9 struct {
	Grid *common.IntGrid
	LowPoints []*common.Coords
	Basins []*Basin
}

type Basin struct {
	LowestPoint *common.Coords
	AllPoints []*common.Coords
}

func (b *Basin) Size() int {
	return len(b.AllPoints)
}

func (p *Puzzle9) FindBasin (c *common.Coords) []*common.Coords {
	ret := []*common.Coords{
		c,
	}
	pointVal, _ := p.Grid.GetCoords(c)
	for _, d := range dirs {
		indir := c.GetCoordsInDir(d, 1)
		neighbor, success := p.Grid.GetCoords(indir)
		if !success {
			continue
		}
		if neighbor > pointVal && neighbor != 9 {
			neiPoints := p.FindBasin(indir)
			for _, np := range neiPoints {
				found := false
				for _, rp := range ret {
					if rp.Col == np.Col && rp.Row == np.Row {
						found = true
					}
				}
				if !found {
					ret = append(ret, np)
				}
			}
		}
	}
	return ret
}

func (p *Puzzle9) FindBasins () {
	for _, l := range p.LowPoints {
		basin := &Basin{
			LowestPoint: l,
			AllPoints: p.FindBasin(l),
		}
		p.Basins = append(p.Basins, basin)
	}
}

func (p *Puzzle9) FindLowPoints () {
	for y, r := range p.Grid.Rows {
		for x, c := range r.Cols {
			coords := common.Coords{
				Row: y,
				Col: x,
			}
			lowest := true
			for _, d := range dirs {
				indir := coords.GetCoordsInDir(d, 1)
				neighbor, success := p.Grid.GetCoords(indir)
				if !success {
					continue
				}
				if neighbor <= c {
					lowest = false
					break
				}
			}
			if lowest {
				p.LowPoints = append(p.LowPoints, &coords)
			}
		}
	}
}

func (p *Puzzle9) ParseHeightmap (input []string) {
	grid := common.NewIntGrid()
	grid.ExtendRows(len(input)-1)
	grid.ExtendCols(len(input[0])-1)
	for y, l := range input {
		e := strings.Split(l, "")
		for x, i := range e {
			height, _ := strconv.Atoi(i)
			grid.Rows[y].Cols[x] = height
		}
	}
	// grid.Print(func(a int) string{
	// 	return fmt.Sprintf("%v", a)
	// })
	p.Grid = grid
}

func (p Puzzle9) Part1(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return nil, err
	}
	output := common.NewSolution(input, "")
	p.ParseHeightmap(i)
	p.FindLowPoints()
	total := 0
	for _, c := range p.LowPoints {
		v, _ := p.Grid.GetCoords(c)
		total += v + 1
	}
	output.Text = fmt.Sprintf("%v", total)
	return output, nil
}

func (p Puzzle9) Part2(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return nil, err
	}
	output := common.NewSolution(input, "")
	p.ParseHeightmap(i)
	p.FindLowPoints()
	p.FindBasins()
	sizes := []int{}
	for _, b := range p.Basins {
		sizes = append(sizes, b.Size())
	}
	top3 := FindTop3(sizes)
	output.Text = fmt.Sprintf("%v", top3[0] * top3[1] * top3[2])
	return output, nil
}

func FindTop3 (ints []int) []int {
	var first, second, third int
	for _, i := range ints {
		if i > first {
			third = second
			second = first
			first = i
		} else if i > second {
			third = second
			second = i
		} else if i > third {
			third = i
		}
	}
	return []int{
		first,
		second,
		third,
	}
}