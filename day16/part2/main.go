package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2017/aoc"
)

const size = 16

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []move {
	line := aoc.ReadAllInput()[0]
	commands := strings.Split(line, ",")
	moves := make([]move, len(commands))
	for i, cmd := range commands {
		var m move
		switch cmd[0] {
		case 's':
			m = spinMove{size: aoc.StrToInt(cmd[1:])}
		case 'x':
			indeces := strings.Split(cmd[1:], "/")
			m = exchangeMove{
				positionA: aoc.StrToInt(indeces[0]),
				positionB: aoc.StrToInt(indeces[1]),
			}
		case 'p':
			m = partnerMove{
				programA: encodeRune(cmd[1]),
				programB: encodeRune(cmd[3]),
			}
		default:
			panic("unexpected command: " + cmd)
		}
		moves[i] = m
	}

	return moves
}

func process(moves []move) string {
	s := newState()

	for i := 0; i < 1_000_000_000; i++ {
		moves[i%len(moves)].run(&s)
	}

	return s.String()
}

func encodeRune(r byte) int {
	return int(r - 'a')
}

func decodeRune(b int) rune {
	return rune(b + 'a')
}

type state struct {
	programs       [size]int
	programToIndex [size]int
	spin           int
}

func newState() state {
	s := state{}
	for i := range s.programs {
		s.programs[i] = i
		s.programToIndex[i] = i
	}

	return s
}

func (s *state) swap(i1, i2 int) {
	i1, i2 = aoc.Mod(i1-s.spin, size), aoc.Mod(i2-s.spin, size)
	s.programs[i1], s.programs[i2] = s.programs[i2], s.programs[i1]
	s.programToIndex[s.programs[i1]], s.programToIndex[s.programs[i2]] = i1, i2
}

func (s *state) rotate(spin int) {
	s.spin = (s.spin + spin) % size
}

func (s *state) index(program int) int {
	return (s.programToIndex[program] + s.spin) % size
}

func (s *state) String() string {
	var b strings.Builder
	for i := 0; i < size; i++ {
		index := aoc.Mod(i-s.spin, size)
		b.WriteRune(decodeRune(s.programs[index]))
	}

	return b.String()
}

type move interface {
	run(*state)
}

type spinMove struct {
	size int
}

func (m spinMove) run(s *state) {
	s.rotate(m.size)
}

type exchangeMove struct {
	positionA int
	positionB int
}

func (m exchangeMove) run(s *state) {
	s.swap(m.positionA, m.positionB)
}

type partnerMove struct {
	programA int
	programB int
}

func (m partnerMove) run(s *state) {
	s.swap(s.index(m.programA), s.index(m.programB))
}
