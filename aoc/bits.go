package aoc

func CountBits[T IntNumber](value T) int {
	bits := 0
	for value != 0 {
		value &= value - 1
		bits++
	}

	return bits
}
