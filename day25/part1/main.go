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

func read() (steps int, states map[rune]state, startState rune) {
	lines := aoc.ReadAllInput()
	startState = rune(lines[0][15])
	steps = aoc.StrToInts(lines[1])[0]
	states = map[rune]state{}
	lines = lines[3:]
	for len(lines) > 0 {
		stateID := rune(lines[0][9])
		lines = lines[1:]
		var s state
		for i := 0; i < 2; i++ {
			conditionValue := aoc.StrToInts(lines[0])[0]
			newValue := aoc.StrToInts(lines[1])[0]
			offset := 1
			if strings.HasSuffix(lines[2], "left.") {
				offset = -1
			}
			nextState := rune(lines[3][26])
			s.actions[conditionValue] = action{
				newValue:  newValue,
				offset:    offset,
				nextState: nextState,
			}

			lines = lines[4:]
		}

		states[stateID] = s

		lines = lines[1:]
	}

	return steps, states, startState
}

func process(steps int, states map[rune]state, startState rune) int {
	currentState := startState
	tape := map[int]int{}
	position := 0
	for i := 0; i < steps; i++ {
		s := states[currentState]
		act := s.actions[tape[position]]
		tape[position] = act.newValue
		position += act.offset
		currentState = act.nextState
	}

	count := 0
	for _, v := range tape {
		if v == 1 {
			count++
		}
	}

	return count
}

type state struct {
	actions [2]action
}

type action struct {
	newValue  int
	offset    int
	nextState rune
}
