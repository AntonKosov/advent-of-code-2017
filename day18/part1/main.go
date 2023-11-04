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
			cmd = rcvCommand{arg: argFunc(parts[1])}
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
	s := newState()
	for {
		cmd := commands[s.position]
		if _, ok := cmd.(rcvCommand); ok && s.lastPlayed != 0 {
			return s.lastPlayed
		}
		cmd.run(&s)
	}
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

type state struct {
	position   int
	registers  map[rune]int
	lastPlayed int
}

func newState() state {
	return state{
		registers: map[rune]int{},
	}
}

type command interface {
	run(*state)
}

type sndCommand struct {
	arg argument
}

func (c sndCommand) run(s *state) {
	s.lastPlayed = c.arg(*s)
	s.position++
}

type setCommand struct {
	register rune
	arg      argument
}

func (c setCommand) run(s *state) {
	s.registers[c.register] = c.arg(*s)
	s.position++
}

type addCommand struct {
	register rune
	arg      argument
}

func (c addCommand) run(s *state) {
	s.registers[c.register] += c.arg(*s)
	s.position++
}

type mulCommand struct {
	register rune
	arg      argument
}

func (c mulCommand) run(s *state) {
	s.registers[c.register] *= c.arg(*s)
	s.position++
}

type modCommand struct {
	register rune
	arg      argument
}

func (c modCommand) run(s *state) {
	s.registers[c.register] = aoc.Mod(s.registers[c.register], c.arg(*s))
	s.position++
}

type rcvCommand struct {
	arg argument
}

func (c rcvCommand) run(s *state) {
	v := c.arg(*s)
	if v != 0 {
		s.lastPlayed = v
	}
	s.position++
}

type jgzCommand struct {
	arg1 argument
	arg2 argument
}

func (c jgzCommand) run(s *state) {
	if c.arg1(*s) > 0 {
		s.position += c.arg2(*s)
		return
	}
	s.position++
}
