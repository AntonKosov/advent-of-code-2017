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
	data = append(data, data[0])
	count := 0
	for i := 0; i < len(data)-1; i++ {
		value := data[i]
		if value == data[i+1] {
			count += int(value - '0')
		}
	}

	return count
}
