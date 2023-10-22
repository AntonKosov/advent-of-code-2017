package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/AntonKosov/advent-of-code-2017/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []string {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]

	return lines
}

func process(passphrases []string) int {
	count := 0
	for _, passphrase := range passphrases {
		words := strings.Split(passphrase, " ")
		usedWords := make(map[string]bool, len(words))
		valid := true
		for _, word := range words {
			letters := []byte(word)
			slices.Sort(letters)
			word = string(letters)
			if usedWords[word] {
				valid = false
				break
			}
			usedWords[word] = true
		}
		if valid {
			count++
		}
	}

	return count
}
