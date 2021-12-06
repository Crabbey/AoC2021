package common

import (
	"os"
	"fmt"
	"bufio"
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
	fmt.Printf("Puzzle %v Part %v Solution: %v\n", a.Input.Puzzle, a.Input.Part, a.Text)
}

func (a *AoCInput) Default(def string) {
	if a.InputFile == "" {
		a.InputFile = def
	}
}

func (a *AoCInput) Read() ([]string, error) {
	fmt.Println("Reading from " + a.Path + "/" +a.InputFile)
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