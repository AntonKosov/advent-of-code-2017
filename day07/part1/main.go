package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2017/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []node {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]
	var graph []node
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		if len(parts) != 2 {
			continue
		}
		graph = append(graph, node{
			parent:   strings.Split(parts[0], " ")[0],
			children: strings.Split(parts[1], ", "),
		})
	}

	return graph
}

func process(graph []node) string {
	children := map[string]bool{}
	for _, n := range graph {
		for _, child := range n.children {
			children[child] = true
		}
	}

	for _, n := range graph {
		if !children[n.parent] {
			return n.parent
		}
	}

	panic("incorrect data")
}

type node struct {
	parent   string
	children []string
}
