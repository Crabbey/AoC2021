package puzzle8

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

type Puzzle8 struct {
	Entries []*Entry
}

type Entry struct {
	First []*Sevenseg
	Second []*Sevenseg
	KnownDigits map[string]*Sevenseg
}

func (e *Entry) DecodeState() {
	for _, s := range e.First {
		t := s.Total()
		switch t {
		case 2: // 1
			e.SetKnownDigits("1", s)
		case 4: // 4
			e.SetKnownDigits("4", s)
		case 3: // 7
			e.SetKnownDigits("7", s)
		case 7: // 8
			e.SetKnownDigits("8", s)
		default:
			continue
		}
	}
	for _, s := range e.First {
		cmp := e.KnownDigits["7"].Compare(s)
		if len(cmp) == 2 {
			e.SetKnownDigits("3", s)
			break
		}
	}
	for _, s := range e.First {
		cmp := e.KnownDigits["3"].Compare(s)
		if len(cmp) == 1 {
			e.SetKnownDigits("9", s)
			break
		}
	}
	for _, s := range e.First {
		cmp := e.KnownDigits["9"].Compare(s)
		if len(cmp) == 1 && s != e.KnownDigits["3"] && s != e.KnownDigits["8"] {
			e.SetKnownDigits("5", s)
			break
		}
	}
	for _, s := range e.First {
		cmp := e.KnownDigits["9"].Compare(s)
		if len(cmp) == 1 && s != e.KnownDigits["3"] && s != e.KnownDigits["8"] {
			e.SetKnownDigits("5", s)
			break
		}
	}
	for _, s := range e.First {
		cmp := e.KnownDigits["5"].Compare(s)
		if len(cmp) == 1 && s != e.KnownDigits["9"]  {
			e.SetKnownDigits("6", s)
			break
		}
	}
	for _, s := range e.First {
		cmpa := e.KnownDigits["8"].Compare(s)
		cmpb := e.KnownDigits["1"].Compare(s)
		if len(cmpa) == 1 && len(cmpb) == 4 && e.KnownDigits["9"] != s {
			e.SetKnownDigits("0", s)
			break
		}
	}
	for _, s := range e.First {
		if s.Total() == 0 {
			continue
		}
		found := false
		for _, b := range e.KnownDigits {
			if s == b {
				found = true
				break
			}
		}
		if !found {
			e.SetKnownDigits("2", s)
		}
	}
}

func (e *Entry) SetKnownDigits(value string, s *Sevenseg) {
	e.KnownDigits[value] = s
}

func (e *Entry) DecodeEntry() int {
	num := ""
	for _, s := range e.Second {
		for digit, g := range e.KnownDigits {
			if len(s.Compare(g)) == 0 {
				num += digit
				break
			}
		}
	}
	val, _ := strconv.Atoi(num)
	return val
}

type Sevenseg struct {
	A bool
	B bool
	C bool
	D bool
	E bool
	F bool
	G bool
}

func (s *Sevenseg) Compare(g *Sevenseg) []string {
	diff := []string{}
	if s.A != g.A {
		diff = append(diff, "A")
	}
	if s.B != g.B {
		diff = append(diff, "B")
	}
	if s.C != g.C {
		diff = append(diff, "C")
	}
	if s.D != g.D {
		diff = append(diff, "D")
	}
	if s.E != g.E {
		diff = append(diff, "E")
	}
	if s.F != g.F {
		diff = append(diff, "F")
	}
	if s.G != g.G {
		diff = append(diff, "G")
	}
	return diff
}

func (s *Sevenseg) Print() string {
	line := ""
	if s.A {
		line += "A"
	}
	if s.B {
		line += "B"
	}
	if s.C {
		line += "C"
	}
	if s.D {
		line += "D"
	}
	if s.E {
		line += "E"
	}
	if s.F {
		line += "F"
	}
	if s.G {
		line += "G"
	}
	return line
}

func (s *Sevenseg) Total() int {
	tot := 0
	if s.A {
		tot++
	}
	if s.B {
		tot++
	}
	if s.C {
		tot++
	}
	if s.D {
		tot++
	}
	if s.E {
		tot++
	}
	if s.F {
		tot++
	}
	if s.G {
		tot++
	}
	return tot
}

func (s *Sevenseg) Parse (i string) {
	l := strings.Split(i, "")
	for _, e := range l {
		switch e {
			case "a":
				s.A = true
			case "b":
				s.B = true
			case "c":
				s.C = true
			case "d":
				s.D = true
			case "e":
				s.E = true
			case "f":
				s.F = true
			case "g":
				s.G = true
		}
	}
}

func (p *Puzzle8) Parse(i []string) {
	for _, l := range i {
		entry := &Entry{
			KnownDigits: make(map[string]*Sevenseg),
		}
		sections := strings.Split(l, "|")
		for a, s := range sections {
			sevensegs := strings.Split(s, " ")
			for _, e := range sevensegs {
				sevenseg := &Sevenseg{}
				sevenseg.Parse(e)
				switch a {
				case 0:
					entry.First = append(entry.First, sevenseg)
				case 1:
					entry.Second = append(entry.Second, sevenseg)
				}
			}
		}
		p.Entries = append(p.Entries, entry)
	}
}

func (p *Puzzle8) CountDistinctNums() int {
	total := 0
	for _, e := range p.Entries {
		for _, s := range e.Second {
			t := s.Total()
			switch t {
			case 2: // 1
				total++
			case 4: // 4
				total++
			case 3: // 7
				total++
			case 7: // 8
				total++
			default:
				continue
			}
		}
	}
	return total
}

func (p Puzzle8) Part1(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return nil, err
	}
	output := common.NewSolution(input, "")
	p.Parse(i)
	tot := p.CountDistinctNums()
	output.Text = fmt.Sprintf("%v", tot)
	return output, nil
}

func (p Puzzle8) Part2(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return nil, err
	}
	output := common.NewSolution(input, "")
	p.Parse(i)
	total := 0
	for _, e := range p.Entries {
		e.DecodeState()
		x := e.DecodeEntry()
		total += x
	}
	output.Text = fmt.Sprintf("%v", total)
	return output, nil
}