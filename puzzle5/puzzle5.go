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
	Vents *Vents
}

type Vents struct {
	Grid *common.IntGrid
}

type Entry struct {
	line string
	setA []string
	setB []string
	setAInt []int
	setBInt []int
}

func (p *Puzzle5) ParseEntry(input string) *Entry {
	a := strings.Split(input, " -> ")
	one := strings.Split(a[0], ",")
	two := strings.Split(a[1], ",")
	ret := Entry{
		line: input,
		setA: one,
		setB: two,
		setAInt: make([]int, 2),
		setBInt: make([]int, 2),
	}
	for x, l := range ret.setA {
		g, _ := strconv.Atoi(l)
		ret.setAInt[x] = g
	}
	for x, l := range ret.setB {
		g, _ := strconv.Atoi(l)
		ret.setBInt[x] = g
	}
	return &ret
}

func (p *Puzzle5) LoadVents1(input []string) {
	v := &Vents{}
	p.Vents = v
	v.Grid = common.NewIntGrid()
	for _, l := range input {
		e := p.ParseEntry(l)
		v.Grid.ExtendRows(max(e.setAInt[1], e.setBInt[1]))
		v.Grid.ExtendCols(max(e.setAInt[0], e.setBInt[0]))
		if e.setA[0] == e.setB[0] || e.setA[1] == e.setB[1] {
			p.DrawLine(e)
		}
	}
	// v.Grid.Print(func(in int)string{
	// 	if in == 0 {
	// 		return "."
	// 	}
	// 	return strconv.Itoa(in)
	// })
}


func (p *Puzzle5) DrawLine(e *Entry) {
	xInc := 0
	yInc := 0
	if (e.setAInt[0] < e.setBInt[0]) {
		// A X < B X
		xInc = 1
	} else if (e.setAInt[0] > e.setBInt[0]) {
		// A X > B X
		xInc = -1
	}

	if (e.setAInt[1] < e.setBInt[1]) {
		// A Y < B Y
		yInc = 1
	} else if (e.setAInt[1] > e.setBInt[1]) {
		// A Y > B Y
		yInc = -1
	}
	xcoord := e.setAInt[0]
	ycoord := e.setAInt[1]
	for {
		p.Vents.Grid.Rows[ycoord].Cols[xcoord].Data++
		xcoord += xInc
		ycoord += yInc
		if xInc > 0 && xcoord > e.setBInt[0] {
			break
		}
		if xInc < 0 && xcoord < e.setBInt[0] {
			break
		}
		if yInc > 0 && ycoord > e.setBInt[1] {
			break
		}
		if yInc < 0 && ycoord < e.setBInt[1] {
			break
		}
	}
}

func (p *Puzzle5) LoadVents2(input []string) {
	v := &Vents{}
	p.Vents = v
	v.Grid = common.NewIntGrid()
	for _, l := range input {
		e := p.ParseEntry(l)
		v.Grid.ExtendRows(max(e.setAInt[1], e.setBInt[1]))
		v.Grid.ExtendCols(max(e.setAInt[0], e.setBInt[0]))
		p.DrawLine(e)
	}
	// v.Grid.Print(func(in int)string{
	// 	if in == 0 {
	// 		return "."
	// 	}
	// 	return strconv.Itoa(in)
	// })
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
	for y := 0; y < len(p.Vents.Grid.Rows); y++ {
		for x := 0; x < len(p.Vents.Grid.Rows[0].Cols); x++ {
			if p.Vents.Grid.Rows[y] == nil {
				fmt.Printf("Unknown row %v\n", y)
				continue
			}
			if p.Vents.Grid.Rows[y].Cols[x] == nil {
				fmt.Printf("Unknown col %v\n", x)
				continue
			}
			if p.Vents.Grid.Rows[y].Cols[x].Data > 1 {
				count++
			}
		}
	}

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
	for y := 0; y < len(p.Vents.Grid.Rows); y++ {
		for x := 0; x < len(p.Vents.Grid.Rows[0].Cols); x++ {
			if p.Vents.Grid.Rows[y] == nil {
				fmt.Printf("Unknown row %v\n", y)
				continue
			}
			if p.Vents.Grid.Rows[y].Cols[x] == nil {
				fmt.Printf("Unknown col %v\n", x)
				continue
			}
			if p.Vents.Grid.Rows[y].Cols[x].Data > 1 {
				count++
			}
		}
	}

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