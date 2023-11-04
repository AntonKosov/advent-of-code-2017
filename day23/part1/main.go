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
		case "set":
			cmd = setCommand{register: rune(parts[1][0]), arg: argFunc(parts[2])}
		case "sub":
			cmd = subCommand{register: rune(parts[1][0]), arg: argFunc(parts[2])}
		case "mul":
			cmd = mulCommand{register: rune(parts[1][0]), arg: argFunc(parts[2])}
		case "jnz":
			cmd = jnzCommand{arg1: argFunc(parts[1]), arg2: argFunc(parts[2])}
		default:
			panic("unexpected command: " + c)
		}

		commands[i] = cmd
	}

	return commands
}

func process(commands []command) int {
	s := newState()
	count := 0
	for s.position >= 0 && s.position < len(commands) {
		cmd := commands[s.position]
		if _, ok := cmd.(mulCommand); ok {
			count++
		}
		cmd.run(&s)
	}

	return count
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
	position  int
	registers map[rune]int
}

func newState() state {
	return state{
		registers: map[rune]int{},
	}
}

type command interface {
	run(*state)
}

type setCommand struct {
	register rune
	arg      argument
}

func (c setCommand) run(s *state) {
	s.registers[c.register] = c.arg(*s)
	s.position++
}

type subCommand struct {
	register rune
	arg      argument
}

func (c subCommand) run(s *state) {
	s.registers[c.register] -= c.arg(*s)
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

type jnzCommand struct {
	arg1 argument
	arg2 argument
}

func (c jnzCommand) run(s *state) {
	if c.arg1(*s) != 0 {
		s.position += c.arg2(*s)
		return
	}
	s.position++
}
