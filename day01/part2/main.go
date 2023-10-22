package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2017/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []byte {
	lines := aoc.ReadAllInput()
	return []byte(lines[0])
}

func process(data []byte) int {
	offset := len(data) / 2
	count := 0
	for i, value := range data {
		i2 := (i + offset) % len(data)
		if value == data[i2] {
			count += int(value - '0')
		}
	}

	return count
}
