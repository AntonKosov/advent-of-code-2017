package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2017/aoc"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []particle {
	lines := aoc.ReadAllInput()
	lines = lines[:len(lines)-1]
	particles := make([]particle, len(lines))
	for i, line := range lines {
		nums := aoc.StrToInts(line)
		particles[i] = particle{
			position:     aoc.NewVector3(nums[0], nums[1], nums[2]),
			velocity:     aoc.NewVector3(nums[3], nums[4], nums[5]),
			acceleration: aoc.NewVector3(nums[6], nums[7], nums[8]),
		}
	}

	return particles
}

func process(particles []particle) int {
	seconds := 1_000
	distances := make([]int, len(particles))
	for i, p := range particles {
		pos := p.position.Add(p.velocity.Mul(seconds)).Add(p.acceleration.Mul(seconds * seconds / 2))
		distances[i] = pos.ManhattanLen()
	}

	minIndex := 0
	for i, d := range distances {
		if distances[minIndex] > d {
			minIndex = i
		}
	}

	return minIndex
}

type particle struct {
	position     aoc.Vector3[int]
	velocity     aoc.Vector3[int]
	acceleration aoc.Vector3[int]
}
