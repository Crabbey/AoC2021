package common

import (
	"fmt"
	"strings"
	"strconv"
	"github.com/davecgh/go-spew/spew"
)

var _ = spew.Dump
var _ = fmt.Println
var _ = strconv.Itoa

type VM struct {
	MemSpace []VMInstruction
	Accumulator int
	Position int
	Terminated bool
}

type VMInstruction struct {
	VM *VM
	Command string
	Arguments []string
}

func (i *VMInstruction) ParseArg(argpos int) int {
	x, _ := strconv.Atoi(i.Arguments[argpos])
	return x
}

func (i *VMInstruction) Execute() {
	switch i.Command {
	case "nop":
		i.VM.Position++
	case "jmp":
		i.VM.Position += i.ParseArg(0)
	case "acc":
		i.VM.Accumulator += i.ParseArg(0)
		i.VM.Position++
		// fmt.Printf("ACC += %v, %v\n", i.ParseArg(0), i.VM.Accumulator)
	default:
		panic("Unknown instruction")
	}
}

func (v *VM) Reset() {
	v.Accumulator = 0
	v.Position = 0
}

func (v *VM) LoadMemspace(input []string) {
	v.MemSpace = nil
	for _, line := range input {
		newInstruction := ReadInstruction(line, v)
		v.MemSpace = append(v.MemSpace, newInstruction)
	}
}

func (v *VM) Step() {
	if v.Position >= len(v.MemSpace) {
		// End of program, terminate.
		v.Terminated = true
		return
	}
	v.MemSpace[v.Position].Execute()
}

func ReadInstruction(input string, vm *VM) VMInstruction {
	ret := VMInstruction{
		VM: vm,
	}
	parts := strings.Split(input, " ")
	ret.Command = parts[0]
	ret.Arguments = parts[1:]
	return ret
}

func NewVM(input []string) *VM {
	ret := VM{}
	ret.LoadMemspace(input)
	return &ret
}