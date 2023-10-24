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

func read() *node {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]
	nodes := map[string]*node{}
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		parentPairs := strings.Split(parts[0], " ")
		parentID := parentPairs[0]
		parentNode := nodes[parentID]
		if parentNode == nil {
			parentNode = &node{}
			nodes[parentID] = parentNode
		}
		parentNode.weight = aoc.StrToInt(parentPairs[1][1 : len(parentPairs[1])-1])
		if len(parts) == 1 {
			continue
		}
		childrenIDs := strings.Split(parts[1], ", ")
		for _, childID := range childrenIDs {
			childNode := nodes[childID]
			if childNode == nil {
				childNode = &node{}
				nodes[childID] = childNode
			}
			parentNode.children = append(parentNode.children, childNode)
		}
	}

	children := map[*node]bool{}
	for _, n := range nodes {
		for _, childNode := range n.children {
			children[childNode] = true
		}
	}

	for _, n := range nodes {
		if !children[n] {
			return n
		}
	}

	panic("root not found")
}

func process(root *node) int {
	var parentNode *node
	sumWeights(root, &parentNode)
	if parentNode == nil {
		panic("tree is balanced")
	}

	weightSums := map[int]int{}
	for _, child := range parentNode.children {
		weightSums[child.sum]++
	}
	if len(weightSums) != 2 {
		panic("unexpected number of weights")
	}

	var unbalancedNode, balancedNode *node
	for _, child := range parentNode.children {
		child := child
		if weightSums[child.sum] == 1 {
			unbalancedNode = child
		} else {
			balancedNode = child
		}
	}

	return unbalancedNode.weight - (unbalancedNode.sum - balancedNode.sum)
}

func sumWeights(n *node, unbalancedNode **node) {
	n.sum = n.weight
	if len(n.children) == 0 {
		return
	}
	weights := map[int]struct{}{}
	for _, child := range n.children {
		sumWeights(child, unbalancedNode)
		if *unbalancedNode != nil {
			return
		}
		n.sum += child.sum
		weights[child.sum] = struct{}{}
	}
	if len(weights) != 1 {
		*unbalancedNode = n
	}
}

type node struct {
	weight   int
	sum      int
	children []*node
}
