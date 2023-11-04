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
	iterations := 2017
	buffer := make([]int, iterations+1)
	pos := 0
	for i := 1; i <= iterations; i++ {
		pos = (pos + steps) % i
		if pos < i-1 {
			copy(buffer[pos+2:], buffer[pos+1:])
		}
		pos++
		buffer[pos] = i
	}

	return buffer[(pos+1)%len(buffer)]
}
