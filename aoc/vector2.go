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

func (v Vector2[T]) Mul(scalar T) Vector2[T] {
	return Vector2[T]{X: v.X * scalar, Y: v.Y * scalar}
}

// RotateLeft rotates the vector to the left (left-handed system)
func (v Vector2[T]) RotateLeft() Vector2[T] {
	return NewVector2(v.Y, -v.X)
}

// RotateRight rotates the vector to the right (left-handed system)
func (v Vector2[T]) RotateRight() Vector2[T] {
	return NewVector2(-v.Y, v.X)
}
