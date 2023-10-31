package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2017/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() [2]generator {
	lines := aoc.ReadAllInput()

	return [2]generator{
		{
			prevValue: aoc.StrToInts(lines[0])[0],
			factor:    16807,
		},
		{
			prevValue: aoc.StrToInts(lines[1])[0],
			factor:    48271,
		},
	}
}

func process(generators [2]generator) int {
	count := 0
	mask := 0xffff
	for i := 0; i < 40_000_000; i++ {
		v1, v2 := generators[0].generate(), generators[1].generate()
		if v1&mask == v2&mask {
			count++
		}
	}

	return count
}

type generator struct {
	prevValue int
	factor    int
}

func (g *generator) generate() int {
	g.prevValue = (g.prevValue * g.factor) % 2147483647

	return g.prevValue
}
