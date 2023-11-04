package aoc

type Queue[T any] struct {
	values []T
}

func (q *Queue[T]) Push(v T) {
	q.values = append(q.values, v)
}

func (q *Queue[T]) Pop() T {
	v := q.values[0]
	q.values = q.values[1:]

	return v
}

func (q *Queue[T]) Size() int {
	return len(q.values)
}
