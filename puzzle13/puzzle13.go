package puzzle13

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

type Puzzle13 struct {
	Grid *common.IntGrid
	FoldLines []*FoldLine
}

type FoldLine struct {
	Line string
	IsRow bool
	Value int
}

func (p *Puzzle13) AddFoldLine(i string) {
	parts := strings.Split(i, "=")
	IsRow := false
	if parts[0] == "fold along y" {
		IsRow = true
	}
	val, _ := strconv.Atoi(parts[1])
	fl := &FoldLine{
		Line: i,
		IsRow: IsRow,
		Value: val,
	}
	p.FoldLines = append(p.FoldLines, fl)
}

func (p *Puzzle13) LoadGrid(input []string) {
	maxcoord := &common.Coords{}
	var allcoords []*common.Coords
	pointsdone := false
	for _, l := range input {
		if l == "" {
			pointsdone = true
			continue
		}
		if pointsdone {
			p.AddFoldLine(l)
			continue
		}
		c := common.CoordFromString(l)
		if c.Row > maxcoord.Row {
			maxcoord.Row = c.Row
		}
		if c.Col > maxcoord.Col {
			maxcoord.Col = c.Col
		}
		allcoords = append(allcoords, c)
	}
	p.Grid = common.NewIntGrid()
	p.Grid.ExtendRows(maxcoord.Row)
	p.Grid.ExtendCols(maxcoord.Col)
	for _, c := range allcoords {
		p.Grid.SetCoords(c, 1)
	}
	p.Grid.SetKey(0, " ")
	p.Grid.SetKey(1, "#")
}

func (p *Puzzle13) ParseAnswer() string {
	// var chars []*common.IntGrid
	ret := ""
	charwidth := 5
	for charInMainGrid := 0; true; charInMainGrid += charwidth {
		if _, ok := p.Grid.Rows[0].Cols[charInMainGrid]; !ok {
			break
		}
		g := common.NewIntGrid()
		// g.SetKey(0, " ")
		g.SetKey(1, "#")
		for y, r := range p.Grid.Rows {
			g.ExtendRows(y)
			g.ExtendCols(charwidth)
			for x := 0; x < charwidth; x++ {
				g.SetCoords(&common.Coords{Row: y, Col: x}, r.Cols[charInMainGrid + x])
			}
		}
		ret += g.DecodeChar()
	}
	return ret
}

func (p *Puzzle13) Fold(f *FoldLine) {
	if f.IsRow {
		for y := 0; true; y++ {
			if _, ok := p.Grid.Rows[f.Value + y]; !ok {
				break
			}
			p.Grid.ExtendRows(f.Value - y)
			p.Grid.ExtendCols(len(p.Grid.Rows[0].Cols)-1)
			for t, v := range p.Grid.Rows[f.Value + y].Cols {
				if v > 0 {
					p.Grid.SetCoords(&common.Coords{Row: f.Value - y, Col: t}, 1)
				}
			}
			delete(p.Grid.Rows, f.Value + y)
		}
		return
	}

	for x := 0; true; x++ {
		// p.Grid.ExtendCols(f.Value - y)
		if _, ok := p.Grid.Rows[0].Cols[f.Value + x]; !ok {
			break
		}
		for rid, _ := range p.Grid.Rows {
			if p.Grid.Rows[rid].Cols[f.Value + x] > 0 {
				p.Grid.SetCoords(&common.Coords{Row: rid, Col: f.Value - x}, 1)
			}
			delete(p.Grid.Rows[rid].Cols, f.Value + x)
		}
	}
}

func (p Puzzle13) Part1(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return nil, err
	}
	output := common.NewSolution(input, "")
	p.LoadGrid(i)
	p.Fold(p.FoldLines[0])
	count := 0
	p.Grid.Foreach(func(v int, c *common.Coords) {
		if v == 1 {
			count++
		}
	})
	output.Text = fmt.Sprintf("%v", count)
	return output, nil
}

func (p Puzzle13) Part2(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return nil, err
	}
	output := common.NewSolution(input, "")
	p.LoadGrid(i)
	for _, l := range p.FoldLines {
		p.Fold(l)
	}
	output.Text = fmt.Sprintf("%v", p.ParseAnswer())
	return output, nil
}