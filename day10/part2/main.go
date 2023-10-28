package main

import (
	"encoding/hex"
	"fmt"

	"github.com/AntonKosov/advent-of-code-2017/aoc"
)

const elements = 256

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []byte {
	bytes := []byte(aoc.ReadAllInput()[0])

	return append(bytes, 17, 31, 73, 47, 23)
}

func process(sequence []byte) string {
	data := initialData()

	rotate(sequence, data)
	denseHash := compress(data)

	return hex.EncodeToString(denseHash)
}

func compress(data []byte) []byte {
	hash := make([]byte, 16)
	for i, v := range data {
		hash[i/16] ^= v
	}

	return hash
}

func rotate(sequence, data []byte) {
	var position byte
	var skipSize byte
	for i := 0; i < 64; i++ {
		for _, length := range sequence {
			reverse(data, position, length)
			position += length + skipSize
			skipSize++
		}
	}
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
