package aoc

type Vector2[T IntNumber] struct {
	X T
	Y T
}

func NewVector2[T Number](x, y T) Vector2[T] {
	return Vector2[T]{X: x, Y: y}
}

func (v Vector2[T]) Add(av Vector2[T]) Vector2[T] {
	return Vector2[T]{X: v.X + av.X, Y: v.Y + av.Y}
}
