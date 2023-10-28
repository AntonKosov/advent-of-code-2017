package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2017/aoc"
)

const (
	// elements = 5 // the example

	elements = 256
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []int {
	line := aoc.ReadAllInput()[0]
	parts := strings.Split(line, ",")
	nums := make([]int, len(parts))
	for i, num := range parts {
		nums[i] = aoc.StrToInt(num)
	}

	return nums
}

func process(sequence []int) int {
	data := initialData()
	position := 0
	for skipSize, length := range sequence {
		reverse(data, position, length)
		position = (position + length + skipSize) % len(data)
	}

	return data[0] * data[1]
}

func initialData() []int {
	data := make([]int, elements)
	for i := range data {
		data[i] = i
	}

	return data
}

func reverse(data []int, start, length int) {
	for i := 0; i < length/2; i++ {
		idx1, idx2 := (start+i)%len(data), (start+length-1-i)%len(data)
		data[idx1], data[idx2] = data[idx2], data[idx1]
	}
}
