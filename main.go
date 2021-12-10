package main

import (
	"fmt"
	"sort"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/crabbey/aoc2021/puzzle1"
	"github.com/crabbey/aoc2021/puzzle2"
	"github.com/crabbey/aoc2021/puzzle3"
	"github.com/crabbey/aoc2021/puzzle4"
	"github.com/crabbey/aoc2021/puzzle5"
	"github.com/crabbey/aoc2021/puzzle6"
	"github.com/crabbey/aoc2021/puzzle7"
	"github.com/crabbey/aoc2021/puzzle8"
	"github.com/crabbey/aoc2021/puzzle9"
	"github.com/crabbey/aoc2021/puzzle10"
	"github.com/crabbey/aoc2021/puzzle11"
	"github.com/crabbey/aoc2021/puzzle12"
	"github.com/crabbey/aoc2021/puzzle13"
	"github.com/crabbey/aoc2021/puzzle14"
	"github.com/crabbey/aoc2021/puzzle15"
	"github.com/crabbey/aoc2021/puzzle16"
	"github.com/crabbey/aoc2021/puzzle17"
	"github.com/crabbey/aoc2021/puzzle18"
	"github.com/crabbey/aoc2021/puzzle19"
	"github.com/crabbey/aoc2021/puzzle20"
	"github.com/crabbey/aoc2021/puzzle21"
	"github.com/crabbey/aoc2021/puzzle22"
	"github.com/crabbey/aoc2021/puzzle23"
	"github.com/crabbey/aoc2021/puzzle24"
	"github.com/crabbey/aoc2021/puzzle25"
	"github.com/crabbey/aoc2021/common"

	"github.com/urfave/cli/v2"
)

var _ = spew.Dump

const implemented = 10

func main() {
	app := cli.NewApp()
	app.Name = "AoC2021 Runner"
	app.EnableBashCompletion = true
	app.CommandNotFound = func(context *cli.Context, cmd string) {
		fmt.Printf("ERROR: Unknown command '%s'\n", cmd)
	}

	app.Commands = []*cli.Command{
		&cmdAllPuzzles,
		&cmdSinglePuzzle,
	}

	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

var cmdAllPuzzles = cli.Command{
	Name:  "all",
	Action: func(c *cli.Context) error {
		for i := 1; i <= implemented; i++ {
			puzzleid := strconv.Itoa(i)
			x := CallPuzzle(c, puzzleid)
			if x != nil {
				return x
			}
		}
		return nil
	},
}

var cmdSinglePuzzle = cli.Command{
	Name:  "puzzle",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "file",
			Aliases: []string{"f"},
			Usage:   "Input file for puzzle",
			EnvVars: []string{"file"},
			Value:   "",
		},
		&cli.BoolFlag{
			Name:    "example",
			Aliases: []string{"e"},
			Usage:   "Use example input",
			Value:   false,
		},
	},
	Action: func(c *cli.Context) error {
		puzzleid := c.Args().Get(0)
		var partid string
		if strings.Contains(puzzleid, ".") {
			puzzlePrompt := strings.Split(puzzleid, ".")
			puzzleid = puzzlePrompt[0]
			partid = puzzlePrompt[1]
			solution, err := CallPuzzlePart(c, puzzleid, partid)
			if err != nil {
				return err
			}
			solution.Print()
			return nil
		}
		CallPuzzle(c, puzzleid)
		return nil
	},
}

func GetInput(c *cli.Context, puzzleid, partid string) common.AoCInput {
	iname := c.String("file")
	if c.Bool("example") {
		iname = "example.txt"
	}
	ret := common.AoCInput{
		Path: "puzzle"+puzzleid,
		InputFile: iname,
		Puzzle: puzzleid,
		Part: partid,
	}
	if iname == "" {
		ret.InputFile = "input.txt"
	}
	return ret
}

func CallPuzzlePart(c *cli.Context, puzzleid string, partid string) (*common.AoCSolution, error) {
	start := time.Now()
	input := GetInput(c, puzzleid, partid)
	var puzzle common.AoCPuzzle
	switch puzzleid {
	case "1":
		puzzle = puzzle1.Puzzle1{}
	case "2":
		puzzle = puzzle2.Puzzle2{}
	case "3":
		puzzle = puzzle3.Puzzle3{}
	case "4":
		puzzle = puzzle4.Puzzle4{}
	case "5":
		puzzle = puzzle5.Puzzle5{}
	case "6":
		puzzle = puzzle6.Puzzle6{}
	case "7":
		puzzle = puzzle7.Puzzle7{}
	case "8":
		puzzle = puzzle8.Puzzle8{}
	case "9":
		puzzle = puzzle9.Puzzle9{}
	case "10":
		puzzle = puzzle10.Puzzle10{}
	case "11":
		puzzle = puzzle11.Puzzle11{}
	case "12":
		puzzle = puzzle12.Puzzle12{}
	case "13":
		puzzle = puzzle13.Puzzle13{}
	case "14":
		puzzle = puzzle14.Puzzle14{}
	case "15":
		puzzle = puzzle15.Puzzle15{}
	case "16":
		puzzle = puzzle16.Puzzle16{}
	case "17":
		puzzle = puzzle17.Puzzle17{}
	case "18":
		puzzle = puzzle18.Puzzle18{}
	case "19":
		puzzle = puzzle19.Puzzle19{}
	case "20":
		puzzle = puzzle20.Puzzle20{}
	case "21":
		puzzle = puzzle21.Puzzle21{}
	case "22":
		puzzle = puzzle22.Puzzle22{}
	case "23":
		puzzle = puzzle23.Puzzle23{}
	case "24":
		puzzle = puzzle24.Puzzle24{}
	case "25":
		puzzle = puzzle25.Puzzle25{}
	default:
		return nil, fmt.Errorf("Unknown puzzle %v", puzzleid)
	}
	var ret *common.AoCSolution
	var err error
	switch partid {
	case "1":
		ret, err = puzzle.Part1(input)
	case "2":
		ret, err = puzzle.Part2(input)
	default:
		return nil, fmt.Errorf("Unknown part id %v", partid)
	}
	ret.Elapsed = time.Since(start)
	return ret, err
}

func CallPuzzle(c *cli.Context, puzzleid string) error {
	parts := []string{"1", "2"}
	for _, x := range parts {
		solution, err := CallPuzzlePart(c, puzzleid, x)
		if err != nil {
			return err
		}
		solution.PrintFancy()		
	}
	return nil
}