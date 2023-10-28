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

func read() []string {
	lines := aoc.ReadAllInput()

	return strings.Split(lines[0], ",")
}

func process(directions []string) int {
	offsets := map[string]aoc.Vector2[int]{
		"n":  aoc.NewVector2(0, -2),
		"ne": aoc.NewVector2(1, -1),
		"se": aoc.NewVector2(1, 1),
		"s":  aoc.NewVector2(0, 2),
		"sw": aoc.NewVector2(-1, 1),
		"nw": aoc.NewVector2(-1, -1),
	}

	pos := aoc.NewVector2[int](0, 0)
	for _, dir := range directions {
		pos = pos.Add(offsets[dir])
	}

	x, y := aoc.Abs(pos.X), aoc.Abs(pos.Y)

	return x + y/2 - x/2
}
