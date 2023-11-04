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

func read() map[pattern]pattern {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]
	patterns := make(map[pattern]pattern, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, " => ")
		patterns[newPattern(parts[0])] = newPattern(parts[1])
	}

	return patterns
}

func process(patterns map[pattern]pattern) int {
	addMissing(patterns)

	pic := newPicture()
	for i := 0; i < 18; i++ {
		pic = expand(pic, patterns)
	}

	return pic.countTurnedOn()
}

func expand(pic picture, patterns map[pattern]pattern) picture {
	from, to := 3, 4
	if pic.size()%2 == 0 {
		from, to = 2, 3
	}
	newPic := newPictureN(pic.size() * to / from)
	for row := 0; row < pic.size(); row += from {
		for col := 0; col < pic.size(); col += from {
			fromPat := pic.readPattern(row, col, from)
			toPat, ok := patterns[fromPat]
			if !ok {
				panic("pattern not found")
			}
			newPic.writePattern(row*to/from, col*to/from, toPat)
		}
	}

	return newPic
}

func addMissing(patternsMap map[pattern]pattern) {
	patterns := aoc.MapKeys(patternsMap)
	for _, p := range patterns {
		to := patternsMap[p]
		add := func(from pattern) {
			if _, ok := patternsMap[from]; !ok {
				patternsMap[from] = to
			}
		}
		for i := 0; i < 4; i++ {
			add(p)
			add(p.flipH())
			add(p.flipV())
			p = p.rotate()
		}
	}
}

type pattern string

func newPattern(p string) pattern {
	return pattern(p)
}

func newPatternFromArray(data [][]bool) pattern {
	var sb strings.Builder
	for _, row := range data {
		if sb.Len() > 0 {
			sb.WriteRune(rowsDelimiter)
		}
		for _, v := range row {
			if v {
				sb.WriteRune(lightOn)
			} else {
				sb.WriteRune(lightOff)
			}
		}
	}

	return pattern(sb.String())
}

func (p pattern) decode() [][]bool {
	rows := strings.Split(string(p), string(rowsDelimiter))
	decoded := make([][]bool, len(rows))
	for r, row := range rows {
		decoded[r] = make([]bool, len(row))
		for c, v := range row {
			decoded[r][c] = v == lightOn
		}
	}

	return decoded
}

func (p pattern) rotate() pattern {
	pat := p.decode()
	if len(pat) == 2 {
		pat[0][0], pat[0][1], pat[1][0], pat[1][1] = pat[1][0], pat[0][0], pat[1][1], pat[0][1]
	} else {
		pat[0][0], pat[0][2], pat[2][0], pat[2][2] = pat[2][0], pat[0][0], pat[2][2], pat[0][2]
		pat[0][1], pat[1][2], pat[2][1], pat[1][0] = pat[1][0], pat[0][1], pat[1][2], pat[2][1]
	}

	return newPatternFromArray(pat)
}

func (p pattern) flipV() pattern {
	pat := p.decode()
	for r, row := range pat {
		for c := 0; c < len(row)/2; c++ {
			c2 := len(row) - c - 1
			pat[r][c], pat[r][c2] = pat[r][c2], pat[r][c]
		}
	}

	return newPatternFromArray(pat)
}

func (p pattern) flipH() pattern {
	pat := p.decode()
	for r, row := range pat {
		r2 := len(pat) - r - 1
		for c := 0; c < len(row)/2; c++ {
			pat[r][c], pat[r2][c] = pat[r2][c], pat[r][c]
		}
	}

	return newPatternFromArray(pat)
}

type picture struct {
	pixels [][]bool
}

func newPicture() picture {
	return picture{
		pixels: [][]bool{
			{false, true, false},
			{false, false, true},
			{true, true, true},
		},
	}
}

func newPictureN(size int) picture {
	p := picture{pixels: make([][]bool, size)}
	for i := range p.pixels {
		p.pixels[i] = make([]bool, size)
	}

	return p
}

func (p picture) countTurnedOn() int {
	count := 0
	for _, row := range p.pixels {
		for _, v := range row {
			if v {
				count++
			}
		}
	}

	return count
}

func (p picture) size() int {
	return len(p.pixels)
}

func (p picture) readPattern(row, col, size int) pattern {
	pat := make([][]bool, size)
	for r := range pat {
		pat[r] = make([]bool, size)
		for c := 0; c < size; c++ {
			pat[r][c] = p.pixels[row+r][col+c]
		}
	}

	return newPatternFromArray(pat)
}

func (p picture) writePattern(row, col int, pat pattern) {
	decPattern := pat.decode()
	for r, patRow := range decPattern {
		for c, v := range patRow {
			p.pixels[row+r][col+c] = v
		}
	}
}

const (
	rowsDelimiter = '/'
	lightOn       = '#'
	lightOff      = '.'
)
