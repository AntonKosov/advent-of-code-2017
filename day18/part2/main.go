package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/AntonKosov/advent-of-code-2017/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []command {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]
	commands := make([]command, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		var cmd command
		switch c := parts[0]; c {
		case "snd":
			cmd = sndCommand{arg: argFunc(parts[1])}
		case "set":
			cmd = setCommand{register: rune(parts[1][0]), arg: argFunc(parts[2])}
		case "add":
			cmd = addCommand{register: rune(parts[1][0]), arg: argFunc(parts[2])}
		case "mul":
			cmd = mulCommand{register: rune(parts[1][0]), arg: argFunc(parts[2])}
		case "mod":
			cmd = modCommand{register: rune(parts[1][0]), arg: argFunc(parts[2])}
		case "rcv":
			cmd = rcvCommand{register: rune(parts[1][0])}
		case "jgz":
			cmd = jgzCommand{arg1: argFunc(parts[1]), arg2: argFunc(parts[2])}
		default:
			panic("unexpected command: " + c)
		}

		commands[i] = cmd
	}

	return commands
}

func process(commands []command) int {
	state0 := newState(0)
	state1 := newState(1)
	for {
		prevPos0, prevPos1 := state0.position, state1.position
		commands[state0.position].run(&state0, &state1.outQueue)
		commands[state1.position].run(&state1, &state0.outQueue)
		if prevPos0 == state0.position && prevPos1 == state1.position {
			break
		}
	}

	return state1.sentValues
}

type argument func(state) int

func argFunc(v string) argument {
	val, err := strconv.Atoi(v)
	if err != nil {
		register := v[0]
		return func(s state) int { return s.registers[rune(register)] }
	}

	return func(s state) int { return val }
}

type inQueue interface {
	Size() int
	Pop() int
}

type state struct {
	sentValues int
	position   int
	registers  map[rune]int
	outQueue   aoc.Queue[int]
}

func newState(programID int) state {
	return state{
		registers: map[rune]int{'p': programID},
	}
}

type command interface {
	run(*state, inQueue)
}

type sndCommand struct {
	arg argument
}

func (c sndCommand) run(s *state, _ inQueue) {
	s.outQueue.Push(c.arg(*s))
	s.sentValues++
	s.position++
}

type setCommand struct {
	register rune
	arg      argument
}

func (c setCommand) run(s *state, _ inQueue) {
	s.registers[c.register] = c.arg(*s)
	s.position++
}

type addCommand struct {
	register rune
	arg      argument
}

func (c addCommand) run(s *state, _ inQueue) {
	s.registers[c.register] += c.arg(*s)
	s.position++
}

type mulCommand struct {
	register rune
	arg      argument
}

func (c mulCommand) run(s *state, _ inQueue) {
	s.registers[c.register] *= c.arg(*s)
	s.position++
}

type modCommand struct {
	register rune
	arg      argument
}

func (c modCommand) run(s *state, _ inQueue) {
	s.registers[c.register] = aoc.Mod(s.registers[c.register], c.arg(*s))
	s.position++
}

type rcvCommand struct {
	register rune
}

func (c rcvCommand) run(s *state, q inQueue) {
	if q.Size() == 0 {
		return
	}

	s.registers[c.register] = q.Pop()
	s.position++
}

type jgzCommand struct {
	arg1 argument
	arg2 argument
}

func (c jgzCommand) run(s *state, _ inQueue) {
	if c.arg1(*s) > 0 {
		s.position += c.arg2(*s)
		return
	}
	s.position++
}
