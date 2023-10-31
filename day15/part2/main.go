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
			mult:      4,
		},
		{
			prevValue: aoc.StrToInts(lines[1])[0],
			factor:    48271,
			mult:      8,
		},
	}
}

func process(generators [2]generator) int {
	count := 0
	mask := 0xffff
	gen0, gen1 := generators[0], generators[1]
	for i := 0; i < 5_000_000; {
		gen0.Generate()
		gen1.Generate()
		if gen0.Empty() || gen1.Empty() {
			continue
		}
		if gen0.Pop()&mask == gen1.Pop()&mask {
			count++
		}
		i++
	}

	return count
}

type generator struct {
	prevValue int
	factor    int
	mult      int
	queue     []int
}

func (g *generator) Generate() {
	g.prevValue = (g.prevValue * g.factor) % 2147483647
	if g.prevValue%g.mult == 0 {
		g.queue = append(g.queue, g.prevValue)
	}
}

func (g *generator) Empty() bool {
	return len(g.queue) == 0
}

func (g *generator) Pop() int {
	v := g.queue[0]
	g.queue = g.queue[1:]

	return v
}
