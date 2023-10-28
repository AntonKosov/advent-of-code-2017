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

func read() []byte {
	line := aoc.ReadAllInput()[0]
	parts := strings.Split(line, ",")
	nums := make([]byte, len(parts))
	for i, num := range parts {
		nums[i] = byte(aoc.StrToInt(num))
	}

	return nums
}

func process(sequence []byte) int {
	data := initialData()
	var position byte
	var skipSize byte
	for _, length := range sequence {
		reverse(data, position, length)
		position += length + skipSize
		skipSize++
	}

	return int(data[0]) * int(data[1])
}

func initialData() []byte {
	data := make([]byte, elements)
	for i := range data {
		data[i] = byte(i)
	}

	return data
}

func reverse(data []byte, start, length byte) {
	for i := byte(0); i < length/2; i++ {
		idx1, idx2 := start+i, start+length-1-i
		data[idx1], data[idx2] = data[idx2], data[idx1]
	}
}
