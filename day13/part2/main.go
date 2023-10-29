package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2017/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []layer {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]
	layers := make([]layer, len(lines))
	for i, line := range lines {
		values := aoc.StrToInts(line)
		layers[i] = layer{depth: values[0], size: values[1]}
	}

	return layers
}

func process(layers []layer) int {
	for delay := 0; ; delay++ {
		if pass(layers, delay) {
			return delay
		}
	}
}

func pass(layers []layer, delay int) bool {
	for _, l := range layers {
		steps := (l.size - 1) * 2
		pos := (steps + l.depth + delay) % steps
		if pos == 0 {
			return false
		}
	}

	return true
}

type layer struct {
	depth int
	size  int
}
