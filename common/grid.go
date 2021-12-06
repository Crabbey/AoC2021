package common

import (
	"fmt"
)


type Grid struct {
	Rows map[int]*Row
}

type Row struct {
	Cols map[int]*Col
}

type Col struct {
	Data string
}

type IntGrid struct {
	Rows map[int]*IntRow
}

type IntRow struct {
	Cols map[int]*IntCol
}

type IntCol struct {
	Data int
}

func NewGrid() *Grid {
	return &Grid{
		Rows: make(map[int]*Row),
	}
}

func (g *Grid) ExtendRows(pos int) {
	for x := 0; x <=pos; x++ {
		if g.Rows[x] == nil {
			g.Rows[x] = &Row{
				Cols: make(map[int]*Col),
			}
		}
	}
}

func (g *Grid) ExtendCols(pos int) {
	for _, b := range g.Rows {
		for x := 0; x <=pos; x++ {
			if b.Cols[x] == nil {
				b.Cols[x] = &Col{}
			}
		}

	}
}

func NewIntGrid() *IntGrid {
	return &IntGrid{
		Rows: make(map[int]*IntRow),
	}
}

func (g *IntGrid) ExtendRows(pos int) {
	for x := 0; x <=pos; x++ {
		if g.Rows[x] == nil {
			g.Rows[x] = &IntRow{
				Cols: make(map[int]*IntCol),
			}
		}
	}
}

func (g *IntGrid) ExtendCols(pos int) {
	for _, b := range g.Rows {
		for x := 0; x <=pos; x++ {
			if b.Cols[x] == nil {
				b.Cols[x] = &IntCol{}
			}
		}

	}
}

func (g *Grid) Print(f func(string)string) {
	for y := 0; y < len(g.Rows); y++ {
		for x := 0; x < len(g.Rows[0].Cols); x++ {
			fmt.Print(f(g.Rows[y].Cols[x].Data))
		}
		fmt.Println()
	}
}

func (g *IntGrid) Print(f func(int)string) {
	for y := 0; y < len(g.Rows); y++ {
		for x := 0; x < len(g.Rows[0].Cols); x++ {
			fmt.Print(f(g.Rows[y].Cols[x].Data))
		}
		fmt.Println()
	}
}