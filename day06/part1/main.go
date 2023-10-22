package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2017/aoc"
)

type hash [16]uint8

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []uint8 {
	lines := aoc.ReadAllInput()
	nums := strings.Split(lines[0], "\t")
	banks := make([]uint8, len(nums))
	for i, num := range nums {
		banks[i] = uint8(aoc.StrToInt(num))
	}

	return banks
}

func process(banks []uint8) int {
	configs := map[hash]bool{
		encode(banks): true,
	}
	cycles := 0
	for {
		cycles++
		redistrubute(banks)
		encoded := encode(banks)
		if configs[encoded] {
			return cycles
		}
		configs[encoded] = true
	}
}

func redistrubute(banks []uint8) {
	maxIndex := findMaxIndex(banks)
	blocks := banks[maxIndex]
	banks[maxIndex] = 0
	for i := maxIndex; blocks > 0; {
		i = (i + 1) % len(banks)
		banks[i]++
		blocks--
	}
}

func findMaxIndex(banks []uint8) int {
	maxIndex := 0
	for i, v := range banks {
		if v > banks[maxIndex] {
			maxIndex = i
		}
	}

	return maxIndex
}

func encode(banks []uint8) hash {
	var h hash
	c := make([]uint8, len(h))
	copy(c, banks)

	return hash(c)
}
