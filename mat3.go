package main

type Mat3 struct {
	A, B, C Vec3
}

// NewMat3 creates and returns a new Mat3, given three vectors that compose its columns.
func NewMat3(a, b, c Vec3) Mat3 {
	return Mat3{a, b, c}
}
