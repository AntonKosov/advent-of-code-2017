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
	m := generateMap(key)
	regions := 0
	for y, row := range m {
		for x, v := range row {
			if v {
				regions++
				fill(m, x, y)
			}
		}
	}

	return regions
}

func fill(m [][]bool, x, y int) {
	if x < 0 || y < 0 || y >= len(m) || x >= len(m[y]) || !m[y][x] {
		return
	}

	m[y][x] = false
	fill(m, x, y-1)
	fill(m, x, y+1)
	fill(m, x-1, y)
	fill(m, x+1, y)
}

func generateMap(key string) [][]bool {
	hashes := generateHashes(key)
	m := make([][]bool, len(hashes))
	for i, h := range hashes {
		m[i] = make([]bool, len(h)*4)
		for j, r := range h {
			value, err := strconv.ParseInt(string(r), 16, 8)
			if err != nil {
				panic(err.Error())
			}
			for k := 0; k < 4; k++ {
				m[i][j*4+k] = value&(1<<(3-k)) != 0
			}
		}
	}

	return m
}

func generateHashes(key string) []string {
	hashes := make([]string, 128)
	for i := range hashes {
		rowKey := fmt.Sprintf("%v-%v", key, i)
		hashes[i] = hash.Generate([]byte(rowKey))
	}

	return hashes
}
