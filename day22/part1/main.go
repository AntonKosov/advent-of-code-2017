package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2017/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() (aoc.Vector2[int], map[aoc.Vector2[int]]bool) {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]
	infected := map[aoc.Vector2[int]]bool{}
	for r, line := range lines {
		for c, v := range line {
			if v == '#' {
				infected[aoc.NewVector2(c, r)] = true
			}
		}
	}

	return aoc.NewVector2(len(lines[0])/2, len(lines)/2), infected
}

func process(pos aoc.Vector2[int], infected map[aoc.Vector2[int]]bool) int {
	dir := aoc.NewVector2(0, -1)
	count := 0
	for i := 0; i < 10_000; i++ {
		if infected[pos] {
			dir = dir.RotateRight()
			delete(infected, pos)
		} else {
			dir = dir.RotateLeft()
			infected[pos] = true
			count++
		}
		pos = pos.Add(dir)
	}

	return count
}
