package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2017/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []int {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]
	instructions := make([]int, len(lines))
	for i, line := range lines {
		instructions[i] = aoc.StrToInt(line)
	}

	return instructions
}

func process(instructions []int) int {
	for currentStep, steps := 0, 0; ; {
		steps++
		offset := &instructions[currentStep]
		currentStep += *offset
		if currentStep < 0 || currentStep >= len(instructions) {
			return steps
		}
		if *offset >= 3 {
			*offset--
		} else {
			*offset++
		}
	}
}
