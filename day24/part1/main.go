package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2017/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []node {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]
	nodes := make([]node, len(lines))
	for i, line := range lines {
		ports := aoc.StrToInts(line)
		n := node{port0: ports[0], port1: ports[1]}
		nodes[i] = n
	}

	return nodes
}

func process(nodes []node) int {
	index := buildIndex(nodes)
	maxStrength := 0
	for _, startNodeIndex := range index[0] {
		n := nodes[startNodeIndex]
		lastPort := n.port0
		if lastPort == 0 {
			lastPort = n.port1
		}
		s := state{
			visited:  1 << startNodeIndex,
			strength: n.strength(),
			lastPort: lastPort,
		}
		cache := map[state]struct{}{}
		findMaxStrength(nodes, index, cache, s)
		for s := range cache {
			maxStrength = max(maxStrength, s.strength)
		}
	}

	return maxStrength
}

func findMaxStrength(nodes []node, index map[int][]int, cache map[state]struct{}, currentState state) {
	if _, ok := cache[currentState]; ok {
		return
	}
	cache[currentState] = struct{}{}

	for _, nextNodeIndex := range index[currentState.lastPort] {
		if currentState.visited&(1<<nextNodeIndex) != 0 {
			continue
		}
		n := nodes[nextNodeIndex]
		p0, p1 := n.port0, n.port1
		lastPort := p0
		if lastPort == currentState.lastPort {
			lastPort = p1
		}
		nextState := state{
			visited:  currentState.visited | (1 << nextNodeIndex),
			strength: currentState.strength + n.strength(),
			lastPort: lastPort,
		}
		findMaxStrength(nodes, index, cache, nextState)
	}
}

func buildIndex(nodes []node) map[int][]int {
	index := map[int][]int{}
	for i, n := range nodes {
		p0, p1 := n.port0, n.port1
		index[p0] = append(index[p0], i)
		if p0 != p1 {
			index[p1] = append(index[p1], i)
		}
	}

	return index
}

type node struct {
	port0 int
	port1 int
}

func (n node) strength() int {
	return n.port0 + n.port1
}

type state struct {
	visited  uint64
	strength int
	lastPort int
}
