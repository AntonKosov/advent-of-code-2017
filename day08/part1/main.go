package main

import (
	"fmt"
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

	actionBuilders := map[string]func(register string, value int) action{
		"inc": func(register string, value int) action {
			return func(reg map[string]int) {
				reg[register] += value
			}
		},
		"dec": func(register string, value int) action {
			return func(reg map[string]int) {
				reg[register] -= value
			}
		},
	}

	conditionBuilders := map[string]func(register string, value int) condition{
		">": func(register string, value int) condition {
			return func(reg map[string]int) bool { return reg[register] > value }
		},
		">=": func(register string, value int) condition {
			return func(reg map[string]int) bool { return reg[register] >= value }
		},
		"<": func(register string, value int) condition {
			return func(reg map[string]int) bool { return reg[register] < value }
		},
		"<=": func(register string, value int) condition {
			return func(reg map[string]int) bool { return reg[register] <= value }
		},
		"==": func(register string, value int) condition {
			return func(reg map[string]int) bool { return reg[register] == value }
		},
		"!=": func(register string, value int) condition {
			return func(reg map[string]int) bool { return reg[register] != value }
		},
	}

	commands := make([]command, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		actValue := aoc.StrToInt(parts[2])
		condValue := aoc.StrToInt(parts[6])
		commands[i] = command{
			action:    actionBuilders[parts[1]](parts[0], actValue),
			condition: conditionBuilders[parts[5]](parts[4], condValue),
		}
	}

	return commands
}

func process(commands []command) int {
	registers := map[string]int{}
	for _, command := range commands {
		command.run(registers)
	}

	var maxValue *int
	for _, value := range registers {
		value := value
		if maxValue == nil || *maxValue < value {
			maxValue = &value
		}
	}

	return *maxValue
}

type action func(reg map[string]int)

type condition func(reg map[string]int) bool

type command struct {
	action    action
	condition condition
}

func (c command) run(registers map[string]int) {
	if c.condition(registers) {
		c.action(registers)
	}
}
