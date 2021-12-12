package puzzle12

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

type Puzzle12 struct {
	StartCave *Cave
	EndCave *Cave
	Caves map[string]*Cave
}

type Cave struct {
	IsBig bool
	Name string
	Connections []*Cave
	Lines []string
}

func (c *Cave) EnsureConnection(dest *Cave) {
	if dest == c {
		return
	}
	found := false
	for _, has := range c.Connections {
		if has == dest {
			found = true
			break
		}
	}
	if !found {
		c.Connections = append(c.Connections, dest)
		dest.Connections = append(dest.Connections, c)
	}
}

func (p *Puzzle12) AddCave(name, line string) {
	cave := &Cave{
		IsBig: (strings.ToUpper(name) == name),
		Name: name,
	}
	cave.Lines = append(cave.Lines, line)
	p.Caves[name] = cave
	if name == "start" {
		p.StartCave = cave
	} else if name == "end" {
		p.EndCave = cave
	}
}

func (p *Puzzle12) LoadCaves(input []string) {
	for _, l := range input {
		parts := strings.Split(l, "-")
		for _, part := range parts {
			if t, ok := p.Caves[part]; ok {
				t.Lines = append(t.Lines, l)
			} else {
				p.AddCave(part, l)
			}
		}
	}
	for _, cave := range p.Caves {
		for _, l := range cave.Lines {
			parts := strings.Split(l, "-")
			p.Caves[parts[0]].EnsureConnection(p.Caves[parts[1]])
		}
	}

	// for _, c := range p.Caves {
	// 	for _, conn := range c.Connections {
	// 		fmt.Printf("%v => %v\n", c.Name, conn.Name)
	// 	}
	// }
}

func (c *Cave) SearchForEnd(p *Path, resp chan []*Path) {
	var ret []*Path
	// p.Print(true)
	count := 0
	ch := make(chan []*Path)
	for _, next := range c.Connections {
		if !p.CanRevisit(next) && !next.IsBig {
			// fmt.Printf("Cave %v ignoring path %v because on path ", c.Name, next.Name)
			continue
		}
		newPath := p.Clone()
		newPath.Add(next)
		if next.Name == "end" {
			ret = append(ret, newPath)
			continue
		}
		count++
		go next.SearchForEnd(newPath, ch)
	}
	run := 0
	for {
		if count == 0 {
			// fmt.Printf("Cave %v exiting (%v) from ", c.Name, run)
			break
		}
		search := <- ch
		count--
		run++
		ret = append(ret, search...)
	}
	// fmt.Printf("Cave %v returning %v paths\n", c.Name, len(ret))
	resp <- ret
}

type Path struct {
	Order map[int]*Cave
	SmallVisitCount int
}

func (p *Path) Add(c *Cave) {
	p.Order[len(p.Order)] = c
}

func (p *Path) Clone() *Path {
	ret := &Path{
		SmallVisitCount: p.SmallVisitCount,
		Order: make(map[int]*Cave),
	}
	for a, b := range p.Order {
		ret.Order[a] = b
	}
	return ret
}

func (p *Path) HasVisitedSmallCaveTwice() bool {
	for x, cave := range p.Order {
		if cave.IsBig {
			continue
		}
		for y, cave2 := range p.Order {
			if x == y {
				continue
			}
			if cave2 == cave {
				return true
			}
		}
	}
	return false
}

func (p *Path) CanRevisit(c *Cave) bool {
	if c.Name == "start" {
		return false
	}
	permittedVisits := p.SmallVisitCount
	if p.HasVisitedSmallCaveTwice() {
		permittedVisits = 1
	}
	count := 0
	for _, cave := range p.Order {
		if cave == c {
			count++
		}
	}
	return count < permittedVisits
}

func (p *Path) Print(indent bool) {
	if indent {
		for x := 0; x <len(p.Order); x++ {
			fmt.Print(" ")
		}
	}
	for x := 0; x <len(p.Order); x++ {
		if x > 0 {
			fmt.Print(", ")
		}
		fmt.Print(p.Order[x].Name)
	}
	fmt.Println()
}

func (p *Puzzle12) FindPaths(numVisits int, print bool) int {
	a := Path{
		SmallVisitCount: numVisits,
	}
	path := a.Clone()
	path.Add(p.StartCave)
	ch := make(chan []*Path)
	go p.StartCave.SearchForEnd(path.Clone(), ch)
	allPaths := <- ch
	if print {
		for _, d := range allPaths {
			d.Print(false)
		}
	}
	return len(allPaths)
}

func (p Puzzle12) Part1(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return nil, err
	}
	output := common.NewSolution(input, "")
	p.Caves = make(map[string]*Cave)
	p.LoadCaves(i)
	output.Text = fmt.Sprintf("%v", p.FindPaths(1, false))
	return output, nil
}

func (p Puzzle12) Part2(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return nil, err
	}
	output := common.NewSolution(input, "")
	p.Caves = make(map[string]*Cave)
	p.LoadCaves(i)
	output.Text = fmt.Sprintf("%v", p.FindPaths(2, false))
	return output, nil
}