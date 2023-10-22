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
	sum := 0
	for _, line := range data {
		sum += findValue(line)
	}

	return sum
}

func findValue(line []int) int {
	for i := 0; i < len(line)-1; i++ {
		vi := line[i]
		for j := i + 1; j < len(line); j++ {
			vj := line[j]
			minV, maxV := min(vi, vj), max(vi, vj)
			if maxV%minV == 0 {
				return maxV / minV
			}
		}
	}

	return 0
}
