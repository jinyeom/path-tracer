package main

type Mat3 struct {
	X, Y, Z Vec3
}

// NewMat3 creates and returns a new Mat3, given three vectors that compose its columns.
func NewMat3(x, y, z Vec3) Mat3 {
	return Mat3{x, y, z}
}
