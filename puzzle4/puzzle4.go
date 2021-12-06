package puzzle4

import (
	"strings"
	"strconv"
	"regexp"
	"fmt"
	"github.com/crabbey/aoc2021/common"
	"github.com/davecgh/go-spew/spew"
)

var _ = spew.Dump
var _ = fmt.Println
var _ = strconv.Itoa
var _ = strings.Split

type Puzzle4 struct {
	Numbers map[int]string
	Boards map[int]*Board
}

type Board struct {
	Rows map[int]*Row
	Complete bool
}

type Row struct {
	Columns map[int]*Number
}

type Number struct {
	Value string
	Marked bool
}

func (b *Board) AddRow(line string) {
	space := regexp.MustCompile(`\s+`)
	startSpace := regexp.MustCompile(`^\s+`)
	newRow := Row{
		Columns: make(map[int]*Number),
	}
	c := space.ReplaceAllString(line, " ")
	l := startSpace.ReplaceAllString(c, "")
	parts := strings.Split(l, " ")
	for p, n := range parts {
		newRow.Columns[p] = &Number{
			Value: n,
		}
	}
	b.Rows[len(b.Rows)] = &newRow
}

func (b *Board) Print() {
	for x := 0; x < len(b.Rows); x++ {
		for y := 0; y < len(b.Rows[x].Columns); y++ {
			fmt.Printf("%v ", b.Rows[x].Columns[y].Value)
		}
		fmt.Println()
	}
}

func (b *Board) MarkNumber(num string) {
	for x := 0; x < len(b.Rows); x++ {
		for y := 0; y < len(b.Rows[x].Columns); y++ {
			if b.Rows[x].Columns[y].Value == num {
				n := b.Rows[x].Columns[y]
				n.Marked = true
				b.Rows[x].Columns[y] = n
			}
		}
	}
}

func (b *Board) Answer(c string) string {
	d, _ := strconv.Atoi(c)
	total := 0
	for x := 0; x < len(b.Rows); x++ {
		for y := 0; y < len(b.Rows[x].Columns); y++ {
			if !b.Rows[x].Columns[y].Marked {
				v, _ := strconv.Atoi(b.Rows[x].Columns[y].Value)
				total += v
			}
		}
	}
	return strconv.Itoa(total * d)
}

func (b *Board) Evaluate() bool {
	for x := 0; x < len(b.Rows); x++ {
		rowMarked := 0
		for y := 0; y < len(b.Rows[x].Columns); y++ {
			if b.Rows[x].Columns[y].Marked {
				rowMarked++
			}
		}
		if rowMarked == 5 {
			b.Complete = true
			return true
		}
	}
	for y := 0; y < len(b.Rows[0].Columns); y++ {
		rowMarked := 0
		for x := 0; x < len(b.Rows); x++ {
			if b.Rows[x].Columns[y].Marked {
				rowMarked++
			}
		}
		if rowMarked == 5 {
			b.Complete = true
			return true
		}
	}
	return false
}

func (p *Puzzle4) ParseInput(input []string) {
	p.Numbers = make(map[int]string)
	p.Boards = make(map[int]*Board)
	CurBoard := &Board{
		Rows: make(map[int]*Row),
	}
	BoardNumber := -1
	for i, line := range input {
		if i == 0 {
			parts := strings.Split(line, ",")
			for x, n := range parts {
				p.Numbers[x] = n
			}
			continue
		}
		if line == "" {
			if BoardNumber == -1 {
				BoardNumber++
				continue;
			}
			// Next Board
			p.Boards[BoardNumber] = CurBoard
			CurBoard = &Board{
				Rows: make(map[int]*Row),
			}
			BoardNumber++
			continue
		}
		CurBoard.AddRow(line)
	}
	// Save the last board
	p.Boards[BoardNumber] = CurBoard
}

func (p *Puzzle4) PrintBoards() {
	for x := 0; x < len(p.Boards); x++ {
		fmt.Printf("---------------\nBoard %v\n", x)
		b := p.Boards[x]
		b.Print()
	}
}

func (p *Puzzle4) Bingo1() string {
	for x := 0; x < len(p.Numbers); x++ {
		num := p.Numbers[x]
		for _, b := range p.Boards {
			b.MarkNumber(num)
			if b.Evaluate() {
				return b.Answer(num)
			}
		}
	}
	return ""
}

func (p *Puzzle4) Bingo2() string {
	lastBoard := &Board{}
	var lastWinningNum string
	for x := 0; x < len(p.Numbers); x++ {
		num := p.Numbers[x]
		for _, b := range p.Boards {
			if b.Complete {
				continue
			}
			b.MarkNumber(num)
			if b.Evaluate() {
				lastBoard = b
				lastWinningNum = num
			}
		}
	}
	return lastBoard.Answer(lastWinningNum)
}

func (p Puzzle4) Part1(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return nil, err
	}
	output := common.NewSolution(input, "")
	p.ParseInput(i)

	output.Text = p.Bingo1()
	return output, nil
}

func (p Puzzle4) Part2(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return nil, err
	}
	output := common.NewSolution(input, "")
	p.ParseInput(i)

	output.Text = p.Bingo2()
	return output, nil
}