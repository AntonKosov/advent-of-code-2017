package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2017/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() string {
	return aoc.ReadAllInput()[0]
}

func process(stream string) int {
	score := 0
	openGroups := 0
	states := []state{expectingGroup}
	push := func(s state) { states = append(states, s) }
	pop := func() { states = states[:len(states)-1] }
	for _, r := range stream {
		s := states[len(states)-1]
		switch s {
		case skipNext:
			pop()
		case waitingEndOfGarbage:
			switch r {
			case skipNextRune:
				push(skipNext)
			case closeGarbage:
				pop()
			default:
				// do nothing
			}
		case expectingGroup:
			switch r {
			case openGroup:
				openGroups++
				score += openGroups
				push(expectingGroup)
			case closeGroup:
				openGroups--
				pop()
				push(expectingGroup)
			case openGarbage:
				push(waitingEndOfGarbage)
			case groupsDelimiter:
				// do nothing
			default:
				panic("unexpected input")
			}
		default:
			panic("unexpected state")
		}
	}

	return score
}

type state int

const (
	expectingGroup state = iota
	skipNext
	waitingEndOfGarbage
)

const (
	skipNextRune    = '!'
	openGroup       = '{'
	closeGroup      = '}'
	openGarbage     = '<'
	closeGarbage    = '>'
	groupsDelimiter = ','
)
