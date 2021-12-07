package common

import (
	"os"
	"fmt"
	"bufio"
	"sync"
	"time"
)

type AoCInput struct {
	Path string
	InputFile string
	Puzzle string
	Part string
}

type AoCSolution struct {
	Input AoCInput
	Text string
	DebugStr string
	mutex sync.Mutex
	Elapsed time.Duration
}

type AoCPuzzle interface {
	Part1(input AoCInput) (*AoCSolution, error)
	Part2(input AoCInput) (*AoCSolution, error)
}

func NewSolution(puzzleinput AoCInput, input string) (*AoCSolution) {
	if input == "" {
		input = "No solution found"
	}
	return &AoCSolution{
		Input: puzzleinput,
		Text: input,
	}
}

func (a *AoCSolution) Print() {
	fmt.Printf("Puzzle %v, Part %v, File %v, Solution: %v\n", a.Input.Puzzle, a.Input.Part, a.Input.InputFile, a.Text)
}

func (a *AoCSolution) Debug(in string) {
	a.mutex.Lock()
	a.DebugStr += in
	a.mutex.Unlock()
}

func (a *AoCSolution) PrintFancy() {
	fmt.Printf("-------------------\n")
	fmt.Printf("Puzzle %v.%v\n", a.Input.Puzzle, a.Input.Part)
	fmt.Printf("Filename %v\n", a.Input.InputFile)
	fmt.Printf("Took %s\n", a.Elapsed)
	fmt.Printf("Answer: %v\n", a.Text)

}

func (a *AoCSolution) PrintDebug() {
	fmt.Printf("%v\n", a.DebugStr)
}

func (a *AoCInput) Default(def string) {
	if a.InputFile == "" {
		a.InputFile = def
	}
}

func (a *AoCInput) Read() ([]string, error) {
	file, err := os.Open(a.Path + "/" +a.InputFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}