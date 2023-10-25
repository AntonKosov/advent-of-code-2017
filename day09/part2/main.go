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
	garbage := 0
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
				garbage++
			}
		case expectingGroup:
			switch r {
			case openGroup:
				push(expectingGroup)
			case closeGroup:
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

	return garbage
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
