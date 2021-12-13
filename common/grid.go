package common

import (
	"fmt"
	"sync"
	"strings"
	"strconv"
	"github.com/davecgh/go-spew/spew"
)

var _ = spew.Dump

type IntGrid struct {
	Rows map[int]*IntRow
	StringKey map[int]string
}

type IntRow struct {
	Cols map[int]int
}

func NewIntGrid() *IntGrid {
	return &IntGrid{
		Rows: make(map[int]*IntRow),
		StringKey: make(map[int]string),
	}
}

func (g *IntGrid) Dimensions() (int, int) {
	return len(g.Rows), len(g.Rows[0].Cols)
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

func (g *IntGrid) SetCoords(c *Coords, v int) bool {
	if _, ok := g.Rows[c.Row]; !ok {
		return false
	}
	if _, ok := g.Rows[c.Row].Cols[c.Col]; !ok {
		return false
	}
	// fmt.Printf("Setting %v, %v from %v to %v\n", c.Row, c.Col, g.Rows[c.Row].Cols[c.Col], v)
	g.Rows[c.Row].Cols[c.Col] = v
	// spew.Dump(g.Rows[c.Row])
	// panic("")
	return true
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

func (g *IntGrid) ExtendCols(width int) {
	wg := &sync.WaitGroup{}
	for _, b := range g.Rows {
		wg.Add(1)
		go func(b *IntRow, wg *sync.WaitGroup, width int) {
			for x := len(b.Cols); x <=width; x++ {
				if _, ok := b.Cols[x]; !ok {
					b.Cols[x] = 0
				}
			}
			wg.Done()
		}(b, wg, width)
	}
	wg.Wait()
}

func (g *IntGrid) Printf(f func(int)string) {
	for y := 0; y < len(g.Rows); y++ {
		for x := 0; x < len(g.Rows[0].Cols); x++ {
			fmt.Print(f(g.Rows[y].Cols[x]))
		}
		fmt.Println()
	}

}

func (g *IntGrid) SetKey(k int, val string) {
	g.StringKey[k] = val
}

func (g *IntGrid) Print() {
	for y := 0; y < len(g.Rows); y++ {
		for x := 0; x < len(g.Rows[0].Cols); x++ {
			if v, ok := g.StringKey[g.Rows[y].Cols[x]]; ok {
				fmt.Print(v)
				continue
			}
			fmt.Print(g.Rows[y].Cols[x])
		}
		fmt.Println()
	}
}

func (g *IntGrid) Foreach(f func(int, *Coords)) {
	for y := 0; y < len(g.Rows); y++ {
		for x := 0; x < len(g.Rows[0].Cols); x++ {
			f(g.Rows[y].Cols[x], &Coords{Row: y, Col: x})
		}
	}
}

func (g *IntGrid) WriteForeach(f func(int, *Coords) int) {
	for y := 0; y < len(g.Rows); y++ {
		for x := 0; x < len(g.Rows[0].Cols); x++ {
			g.Rows[y].Cols[x] = f(g.Rows[y].Cols[x], &Coords{Row: y, Col: x})
		}
	}
}

func StrToInt(s string) int {
	a, _ := strconv.Atoi(s)
	return a
}

func (g *IntGrid) LoadFromFile(input []string, f func(string)int) {
	g.ExtendRows(len(input)-1)
	g.ExtendCols(len(input[0])-1)
	for ycoord, line := range input {
		chars := strings.Split(line, "")
		for xcoord, char := range chars {
			val := f(char)
			g.SetKey(val, char)
			g.Rows[ycoord].Cols[xcoord] = val
		}
	}

}