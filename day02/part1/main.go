package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2017/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() [][]int {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]
	data := make([][]int, len(lines))
	for i, line := range lines {
		data[i] = aoc.StrToInts(line)
	}

	return data
}

func process(data [][]int) int {
	checksum := 0
	for _, line := range data {
		minV, maxV := line[0], line[0]
		for _, v := range line {
			minV = min(minV, v)
			maxV = max(maxV, v)
		}
		checksum += maxV - minV
	}

	return checksum
}
