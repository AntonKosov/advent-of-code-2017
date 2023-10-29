package main

import (
	"fmt"
	"strconv"

	"github.com/AntonKosov/advent-of-code-2017/aoc"
	"github.com/AntonKosov/advent-of-code-2017/day10/part2/hash"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() string {
	return aoc.ReadAllInput()[0]
}

func process(key string) int {
	hashes := generateHashes(key)
	used := 0
	for _, hash := range hashes {
		used += countUsed(hash)
	}

	return used
}

func countUsed(hash string) int {
	used := 0
	for _, r := range hash {
		value, err := strconv.ParseInt(string(r), 16, 8)
		if err != nil {
			panic(err.Error())
		}
		used += aoc.CountBits(value)
	}

	return used
}

func generateHashes(key string) []string {
	hashes := make([]string, 128)
	for i := range hashes {
		rowKey := fmt.Sprintf("%v-%v", key, i)
		hashes[i] = hash.Generate([]byte(rowKey))
	}

	return hashes
}
