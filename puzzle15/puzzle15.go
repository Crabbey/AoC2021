package puzzle15

import (
	"strings"
	"strconv"
	"fmt"
	"github.com/crabbey/aoc2021/common"
	"github.com/davecgh/go-spew/spew"
	"github.com/RyanCarrier/dijkstra"
)

var _ = spew.Dump
var _ = strings.Split
var _ = strconv.Itoa
var _ = fmt.Printf

type Puzzle15 struct {
	Grid *common.IntGrid
	Graph *dijkstra.Graph
}

func (p Puzzle15) Part1(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return nil, err
	}
	output := common.NewSolution(input, "")
	p.Grid = common.NewIntGrid()
	p.Grid.LoadFromFile(i, common.StrToInt)

	dimX, dimY := p.Grid.Dimensions()
	maxcoord := fmt.Sprintf("%v|%v", dimX-1, dimY-1)
	fmt.Sprintf("%v", maxcoord)

	graph, vertices := p.Grid.CreateDijkstra()
	p.Graph = graph

	path, err := p.Graph.Shortest(vertices["0|0"].ID, vertices[maxcoord].ID)
	if err != nil {
		return output, err
	}
	output.Text = fmt.Sprintf("%v", path.Distance)

	return output, nil
}

func (p Puzzle15) Part2(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return nil, err
	}
	output := common.NewSolution(input, "")
	baseGrid := common.NewIntGrid()
	baseGrid.LoadFromFile(i, common.StrToInt)
	baseX, baseY := baseGrid.Dimensions()
	dimX := baseX * 5
	dimY := baseY * 5

	output.Benchmark("BeginGrid")
	bigGrid := common.NewIntGrid()
	bigGrid.ExtendRows(dimY-1)
	bigGrid.ExtendCols(dimX-1)

	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			dist := x + y
			baseGrid.Foreach(func(v int, coords *common.Coords){
				destCoords := &common.Coords {
					Row: coords.Row + (y * baseY),
					Col: coords.Col + (x * baseX),
				}
				value := ((v - 1 + dist) % 9) + 1 // v-1 to offset the 1-9 modulo, +1 at the end
				bigGrid.SetCoords(destCoords, value)
			})
		}
	}
	output.Benchmark("EndGrid")

	maxcoord := fmt.Sprintf("%v|%v", dimX-1, dimY-1)
	fmt.Sprintf("%v", maxcoord)

	graph, vertices := bigGrid.CreateDijkstra()
	output.Benchmark("CreateGraph")

	p.Graph = graph

	path, err := p.Graph.Shortest(vertices["0|0"].ID, vertices[maxcoord].ID)
	if err != nil {
		return output, err
	}
	output.Benchmark("FindPath")
	output.Text = fmt.Sprintf("%v", path.Distance)
	return output, nil
}
