package main

import (
	"fmt"
	"math"

	"github.com/AntonKosov/advent-of-code-2017/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() int {
	lines := aoc.ReadAllInput()

	return aoc.StrToInt(lines[0])
}

func process(data int) int {
	size := getSize(data)
	startValue := (size-2)*(size-2) + 1
	data = startValue + (data-startValue)%(size-1)
	zeroValue := startValue + size/2 - 1

	return size/2 + aoc.Abs(data-zeroValue)
}

func getSize(data int) int {
	size := int(math.Sqrt(float64(data)))
	if size&1 == 0 {
		size++
	}
	if size*size < data {
		size += 2
	}

	return size
}
