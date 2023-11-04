package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2017/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() [][]rune {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-2]
	diagram := make([][]rune, len(lines))
	for i, line := range lines {
		diagram[i] = []rune(line)
	}

	return diagram
}

func process(diagram [][]rune) int {
	dirs := []aoc.Vector2[int]{
		aoc.NewVector2(0, 1),
		aoc.NewVector2(0, -1),
		aoc.NewVector2(1, 0),
		aoc.NewVector2(-1, 0),
	}
	pos := findEnter(diagram)
	prevPos := pos
	prevPos.Y--
	dir := dirs[0]
	for steps := 0; ; steps++ {
		switch r := diagram[pos.Y][pos.X]; r {
		case vertical:
			// nothing
		case horizontal:
			// nothing
		case changeDirection:
			for _, d := range dirs {
				np := pos.Add(d)
				if np != prevPos && np.Y < len(diagram) && diagram[np.Y][np.X] != exit {
					dir = d
					break
				}
			}
		case exit:
			return steps
		}
		prevPos = pos
		pos = pos.Add(dir)
	}
}

func findEnter(diagram [][]rune) aoc.Vector2[int] {
	for i, v := range diagram[0] {
		if v == vertical {
			return aoc.NewVector2(i, 0)
		}
	}

	panic("enter not found")
}

const (
	vertical        = '|'
	horizontal      = '-'
	changeDirection = '+'
	exit            = ' '
)
