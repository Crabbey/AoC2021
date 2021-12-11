package puzzle11

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

var directions = []string{
	"left",
	"right",
	"up",
	"down",
	"upleft",
	"upright",
	"downleft",
	"downright",
}

type Puzzle11 struct {
	Grid *common.IntGrid
	FlashCount int
	Flashed map[string]*common.Coords
}

func (p *Puzzle11) EvalCoord(noop int, coord *common.Coords) {
	val, _ := p.Grid.GetCoords(coord)
	if _, ok := p.Flashed[coord.UniqueReference()]; !ok && val > 9 {
		p.Flashed[coord.UniqueReference()] = coord
		for _, d := range directions {
			indir := coord.GetCoordsInDir(d, 1)
			value, success := p.Grid.GetCoords(indir)
			if !success {
				continue
			}
			p.Grid.SetCoords(indir, value + 1)
			p.EvalCoord(0, indir)
		}
	}
}

func (p *Puzzle11) Tick() {
	p.Flashed = make(map[string]*common.Coords)
	p.Grid.WriteForeach(func(val int, coord *common.Coords) int {
		value, _ := p.Grid.GetCoords(coord)
		return value + 1
	})
	p.Grid.Foreach(p.EvalCoord)
	for _, coord := range p.Flashed {
		p.FlashCount++
		p.Grid.SetCoords(coord, 0)
	}
}

func (p Puzzle11) Part1(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return nil, err
	}
	output := common.NewSolution(input, "")
	p.Grid = common.NewIntGrid()
	p.Grid.LoadFromFile(i, common.StrToInt)

	p.Flashed = make(map[string]*common.Coords)
	for x := 0; x < 100; x++ {
		p.Tick()
	}
	output.Text = fmt.Sprintf("%v", p.FlashCount)
	return output, nil
}

func (p Puzzle11) Part2(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return nil, err
	}
	output := common.NewSolution(input, "")
	p.Grid = common.NewIntGrid()
	p.Grid.LoadFromFile(i, common.StrToInt)

	p.Flashed = make(map[string]*common.Coords)
	count := 0
	for {
		count++
		p.Tick()
		if len(p.Flashed) == 100 {
			break
		}
	}
	output.Text = fmt.Sprintf("%v", count)
	return output, nil
}

