package puzzle5

import (
	"strings"
	"strconv"
	"fmt"
	"github.com/crabbey/aoc2021/common"
	"github.com/davecgh/go-spew/spew"
)

var _ = spew.Dump
var _ = fmt.Println

type Puzzle5 struct {
	Grid *common.IntGrid
}

type Entry struct {
	line string
	SetA *common.Coords
	SetB *common.Coords
}

func (p *Puzzle5) ParseEntry(input string) *Entry {
	a := strings.Split(input, " -> ")
	one := strings.Split(a[0], ",")
	two := strings.Split(a[1], ",")
	ret := Entry{
		line: input,
		SetA: &common.Coords{},
		SetB: &common.Coords{},
	}
	for x, l := range one {
		switch x {
			case 0:
				ret.SetA.Col, _ = strconv.Atoi(l)
			case 1:
				ret.SetA.Row, _ = strconv.Atoi(l)
		}
	}
	for x, l := range two {
		switch x {
			case 0:
				ret.SetB.Col, _ = strconv.Atoi(l)
			case 1:
				ret.SetB.Row, _ = strconv.Atoi(l)
		}
	}
	return &ret
}

func (p *Puzzle5) LoadVents1(input []string) {
	p.Grid = common.NewIntGrid()
	var entries []*Entry
	var highRow int
	var highCol int
	for _, l := range input {
		e := p.ParseEntry(l)
		entries = append(entries, e)
		highRow = max(highRow, max(e.SetA.Row, e.SetB.Row))
		highCol = max(highCol, max(e.SetA.Col, e.SetB.Col))
	}
	p.Grid.ExtendRows(highRow)
	p.Grid.ExtendCols(highCol)

	for _, e := range entries {
		diff := e.SetA.Diff(e.SetB)
		if diff.Row == 0 || diff.Col == 0 {
			p.DrawLine(e)
		}
	}
	// p.Grid.SetKey(0, ".")
	// p.Grid.Print()
}


func (p *Puzzle5) DrawLine(e *Entry) {
	vector := &common.Coords{}
	if (e.SetA.Col < e.SetB.Col) {
		// A X < B X
		vector.Col = 1
	} else if (e.SetA.Col > e.SetB.Col) {
		// A X > B X
		vector.Col = -1
	}

	if (e.SetA.Row < e.SetB.Row) {
		// A Y < B Y
		vector.Row = 1
	} else if (e.SetA.Row > e.SetB.Row) {
		// A Y > B Y
		vector.Row = -1
	}
	activeCoord := &common.Coords{
		Col: e.SetA.Col,
		Row: e.SetA.Row,
	}
	for {
		p.Grid.Rows[activeCoord.Row].Cols[activeCoord.Col]++
		activeCoord.Translate(vector, 1)
		if vector.Col > 0 && activeCoord.Col > e.SetB.Col {
			break
		} else if vector.Col < 0 && activeCoord.Col < e.SetB.Col {
			break
		} else if vector.Row > 0 && activeCoord.Row > e.SetB.Row {
			break
		} else if vector.Row < 0 && activeCoord.Row < e.SetB.Row {
			break
		}
	}
}

func (p *Puzzle5) LoadVents2(input []string) {
	p.Grid = common.NewIntGrid()
	var entries []*Entry
	var highRow int
	var highCol int
	for _, l := range input {
		e := p.ParseEntry(l)
		entries = append(entries, e)
		highRow = max(highRow, max(e.SetA.Row, e.SetB.Row))
		highCol = max(highCol, max(e.SetA.Col, e.SetB.Col))
	}
	p.Grid.ExtendRows(highRow)
	p.Grid.ExtendCols(highCol)

	for _, e := range entries {
		p.DrawLine(e)
	}
	// p.Grid.SetKey(0, ".")
	// p.Grid.Print()
}

func (p Puzzle5) Part1(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return nil, err
	}
	output := common.NewSolution(input, "")
	p.LoadVents1(i)
	count := 0
	p.Grid.Foreach(func(i int) {
		if i > 1 {
			count++
		}
	})

	output.Text = fmt.Sprintf("%v", count)
	return output, nil
}

func (p Puzzle5) Part2(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return nil, err
	}
	output := common.NewSolution(input, "")
	p.LoadVents2(i)
	count := 0
	p.Grid.Foreach(func(i int) {
		if i > 1 {
			count++
		}
	})


	output.Text = fmt.Sprintf("%v", count)
	return output, nil
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}