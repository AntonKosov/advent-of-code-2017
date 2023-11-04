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
	for i := 0; i < seconds; i++ {
		particles = move(particles)
	}

	return len(particles)
}

func move(particles []particle) []particle {
	positions := make(map[aoc.Vector3[int]]int, len(particles))
	for i := range particles {
		p := &particles[i]
		p.velocity = p.velocity.Add(p.acceleration)
		p.position = p.position.Add(p.velocity)
		positions[p.position]++
	}

	result := make([]particle, 0, len(positions))
	for _, p := range particles {
		if positions[p.position] == 1 {
			result = append(result, p)
		}
	}

	return result
}

type particle struct {
	position     aoc.Vector3[int]
	velocity     aoc.Vector3[int]
	acceleration aoc.Vector3[int]
}
