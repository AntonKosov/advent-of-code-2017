package aoc

type Vector3[T IntNumber] struct {
	X T
	Y T
	Z T
}

func NewVector3[T Number](x, y, z T) Vector3[T] {
	return Vector3[T]{X: x, Y: y, Z: z}
}

func (v Vector3[T]) Add(av Vector3[T]) Vector3[T] {
	return Vector3[T]{X: v.X + av.X, Y: v.Y + av.Y, Z: v.Z + av.Z}
}

func (v Vector3[T]) Mul(scalar T) Vector3[T] {
	return Vector3[T]{X: v.X * scalar, Y: v.Y * scalar, Z: v.Z * scalar}
}

func (v Vector3[T]) ManhattanLen() T {
	return Abs(v.X) + Abs(v.Y) + Abs(v.Z)
}
