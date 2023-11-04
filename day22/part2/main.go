package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2017/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() (aoc.Vector2[int], map[aoc.Vector2[int]]node) {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]
	nodes := map[aoc.Vector2[int]]node{}
	for r, line := range lines {
		for c, v := range line {
			if v == '#' {
				nodes[aoc.NewVector2(c, r)] = infectedNode
			}
		}
	}

	return aoc.NewVector2(len(lines[0])/2, len(lines)/2), nodes
}

func process(pos aoc.Vector2[int], nodes map[aoc.Vector2[int]]node) int {
	dir := aoc.NewVector2(0, -1)
	count := 0
	for i := 0; i < 10_000_000; i++ {
		switch nodes[pos] {
		case cleanNode:
			nodes[pos] = weakenedNode
			dir = dir.RotateLeft()
		case weakenedNode:
			nodes[pos] = infectedNode
			count++
		case infectedNode:
			nodes[pos] = flaggedNode
			dir = dir.RotateRight()
		case flaggedNode:
			delete(nodes, pos)
			dir = dir.Mul(-1)
		}
		pos = pos.Add(dir)
	}

	return count
}

type node byte

const (
	cleanNode    node = 0
	weakenedNode node = 1
	infectedNode node = 2
	flaggedNode  node = 3
)
