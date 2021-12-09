package common

import (
	"fmt"
	"sync"
)


type Coords struct {
	Row int
	Col int
}

type Grid struct {
	Rows map[int]*Row
}

type Row struct {
	Cols map[int]string
}

type IntGrid struct {
	Rows map[int]*IntRow
}

type IntRow struct {
	Cols map[int]int
}

func NewGrid() *Grid {
	return &Grid{
		Rows: make(map[int]*Row),
	}
}

func (c *Coords) Print() (string) {
	return fmt.Sprintf("%v, %v", c.Col, c.Row)
}

func (c *Coords) GetCoordsInDir(dir string, distance int) (*Coords) {
	newCoords := Coords{
		Row: c.Row,
		Col: c.Col,
	}
	switch dir {
		case "left":
			newCoords.Col -= distance
		case "right":
			newCoords.Col += distance
		case "up":
			newCoords.Row -= distance
		case "down":
			newCoords.Row += distance
	}
	return &newCoords
}

func (g *Grid) ExtendRows(pos int) {
	for x := 0; x <=pos; x++ {
		if g.Rows[x] == nil {
			g.Rows[x] = &Row{
				Cols: make(map[int]string),
			}
		}
	}
}

func (g *Grid) ExtendCols(pos int) {
	wg := &sync.WaitGroup{}
	for _, b := range g.Rows {
		wg.Add(1)
		go func(b *Row, wg *sync.WaitGroup, pos int) {
			for x := len(b.Cols); x <=pos; x++ {
				if _, ok := b.Cols[x]; !ok {
					b.Cols[x] = ""
				}
			}
			wg.Done()
		}(b, wg, pos)
	}
	wg.Wait()
}



func NewIntGrid() *IntGrid {
	return &IntGrid{
		Rows: make(map[int]*IntRow),
	}
}

func (g *IntGrid) GetCoords(c *Coords) (int, bool) {
	if _, ok := g.Rows[c.Row]; !ok {
		return 0, false
	}
	if _, ok := g.Rows[c.Row].Cols[c.Col]; !ok {
		return 0, false
	}
	return g.Rows[c.Row].Cols[c.Col], true
}


func (g *IntGrid) ExtendRows(pos int) {
	for x := 0; x <=pos; x++ {
		if g.Rows[x] == nil {
			g.Rows[x] = &IntRow{
				Cols: make(map[int]int),
			}
		}
	}
}

func (g *IntGrid) ExtendCols(pos int) {
	wg := &sync.WaitGroup{}
	for _, b := range g.Rows {
		wg.Add(1)
		go func(b *IntRow, wg *sync.WaitGroup, pos int) {
			for x := len(b.Cols); x <=pos; x++ {
				if _, ok := b.Cols[x]; !ok {
					b.Cols[x] = 0
				}
			}
			wg.Done()
		}(b, wg, pos)
	}
	wg.Wait()
}

func (g *Grid) Print(f func(string)string) {
	for y := 0; y < len(g.Rows); y++ {
		for x := 0; x < len(g.Rows[0].Cols); x++ {
			fmt.Print(f(g.Rows[y].Cols[x]))
		}
		fmt.Println()
	}
}

func (g *IntGrid) Print(f func(int)string) {
	for y := 0; y < len(g.Rows); y++ {
		for x := 0; x < len(g.Rows[0].Cols); x++ {
			fmt.Print(f(g.Rows[y].Cols[x]))
		}
		fmt.Println()
	}
}