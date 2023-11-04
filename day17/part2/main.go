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
	return aoc.StrToInt(aoc.ReadAllInput()[0])
}

func process(steps int) int {
	answer := 0
	for i, pos, zeroPos := 1, 0, 0; i <= 50_000_000; i++ {
		pos = (pos + steps) % i
		switch {
		case pos < zeroPos:
			zeroPos++
		case pos == zeroPos:
			answer = i
		}
		pos++
	}

	return answer
}
