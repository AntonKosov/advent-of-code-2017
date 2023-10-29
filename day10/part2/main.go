package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2017/aoc"
	"github.com/AntonKosov/advent-of-code-2017/day10/part2/hash"
)

func main() {
	answer := hash.Generate(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []byte {
	return []byte(aoc.ReadAllInput()[0])
}
