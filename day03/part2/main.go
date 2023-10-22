package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2017/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() int {
	lines := aoc.ReadAllInput()

	return aoc.StrToInt(lines[0])
}

func process(data int) int {
	values := map[aoc.Vector2[int]]int{
		aoc.NewVector2(0, 0): 1,
	}
	size := 1
	pos := aoc.NewVector2(0, 0)
	offsets := []aoc.Vector2[int]{
		aoc.NewVector2(0, -1),
		aoc.NewVector2(-1, 0),
		aoc.NewVector2(0, 1),
		aoc.NewVector2(1, 0),
	}

	for {
		size += 2
		pos.X++
		pos.Y++
		for _, offset := range offsets {
			for i := 0; i < size-1; i++ {
				pos = pos.Add(offset)
				sum := 0
				for x := pos.X - 1; x <= pos.X+1; x++ {
					for y := pos.Y - 1; y <= pos.Y+1; y++ {
						sum += values[aoc.NewVector2(x, y)]
					}
				}
				if sum > data {
					return sum
				}
				values[pos] = sum
			}
		}
	}
}
