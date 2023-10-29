package aoc

type IntNumber interface {
	~int | ~int64
}

type Number interface {
	IntNumber
}

func Abs[T Number](a T) T {
	if a < 0 {
		return -a
	}

	return a
}
