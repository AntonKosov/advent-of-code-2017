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
	adjList := make([][]int, len(lines))
	for i, line := range lines {
		nums := aoc.StrToInts(line)
		adjList[i] = nums[1:]
	}

	return adjList
}

func process(adjList [][]int) int {
	visited := make([]bool, len(adjList))
	count := 0
	for i, v := range visited {
		if !v {
			visit(adjList, i, visited)
			count++
		}
	}

	return count
}

func visit(adjList [][]int, index int, visited []bool) {
	if visited[index] {
		return
	}

	visited[index] = true

	for _, adjIndex := range adjList[index] {
		visit(adjList, adjIndex, visited)
	}
}
