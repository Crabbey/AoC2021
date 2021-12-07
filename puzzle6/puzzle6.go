package puzzle6

import (
	"strings"
	"strconv"
    "sync"
	"fmt"
	"github.com/crabbey/aoc2021/common"
	"github.com/davecgh/go-spew/spew"
	"time"
)

var _ = spew.Dump
var _ = strings.Split
var _ = strconv.Itoa
var _ = time.Second

type Puzzle6 struct {
	AllFish []*Fish
	Mutex *sync.Mutex
	FishEmulator *FishEmulator
}

// Original solution
type Fish struct {
	p *Puzzle6
	counter int
}

func (f *Fish) Tick() {
	f.counter--
	if f.counter == -1 {
		f.p.NewFish(8)
		f.counter = 6
	}
}

type FishEmulator struct {
	FishCount map[int]int
}

func (f *FishEmulator) Tick() {
	newCount := make(map[int]int)
	newCount[-1] = 0
	newCount[0] = 0
	newCount[6] = 0
	for a, b := range f.FishCount {
		newCount[a-1] = b
	}
	newCount[6] += newCount[-1]
	newCount[8] += newCount[-1]
	delete(newCount, -1)
	f.FishCount = newCount
}

func (f *FishEmulator) Init() {
	f.FishCount = make(map[int]int)
}

func (f *FishEmulator) Total() int {
	total := 0
	for _, b := range f.FishCount {
		total += b
	}
	return total
}

// Defunct - original part 1 solution
func (p *Puzzle6) Tick() {
	var wg sync.WaitGroup
	for _, f := range p.AllFish {
		wg.Add(1)
		go func(f *Fish, wg *sync.WaitGroup) {
			defer wg.Done()
			f.Tick()
		}(f, &wg)
	}
	wg.Wait()
}

// Defunct - original part 1 solution
func (p *Puzzle6) NewFish(counter int) {
	nf := &Fish{
		counter: counter,
		p: p,
	}
	p.Mutex.Lock()
	p.AllFish = append(p.AllFish, nf)
	p.Mutex.Unlock()
}

func (p Puzzle6) Part1(input common.AoCInput) (*common.AoCSolution, error) {
	p.Mutex = &sync.Mutex{}
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return nil, err
	}
	p.FishEmulator = &FishEmulator{}
	p.FishEmulator.Init()
	output := common.NewSolution(input, "")
	fishCounters := strings.Split(i[0], ",")
	for _, c := range fishCounters {
		i, _ := strconv.Atoi(c)
		p.FishEmulator.FishCount[i]++
	}
	for x := 0; x < 80; x++ {
		p.FishEmulator.Tick()
	}
	output.Text = fmt.Sprintf("%v", p.FishEmulator.Total())
	return output, nil
}

func (p Puzzle6) Part2(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return nil, err
	}
	p.FishEmulator = &FishEmulator{}
	p.FishEmulator.Init()
	output := common.NewSolution(input, "")
	fishCounters := strings.Split(i[0], ",")
	for _, c := range fishCounters {
		i, _ := strconv.Atoi(c)
		p.FishEmulator.FishCount[i]++
	}
	for x := 0; x < 256; x++ {
		p.FishEmulator.Tick()
	}
	output.Text = fmt.Sprintf("%v", p.FishEmulator.Total())
	return output, nil
}