package main

import (
	"fmt"
	"math"
)

// Vec3 is a vector of 3 real numbers.
type Vec3 struct {
	X, Y, Z float64
}

// NewVec3 returns a new Vec3.
func NewVec3(x, y, z float64) *Vec3 {
	return &Vec3{x, y, z}
}

// Copy returns a deep copy of the vector
func (v *Vec3) Copy() *Vec3 {
	return NewVec3(v.X, v.Y, v.Z)
}

// String returns the string representation of the vector.
func (v *Vec3) String() string {
	return fmt.Sprintf("vec3(%f, %f, %f)", v.X, v.Y, v.Z)
}

// Add returns the sum of this vector and another vector.
func (v *Vec3) Add(v1 *Vec3) *Vec3 {
	return &Vec3{v.X + v1.X, v.Y + v1.Y, v.Z + v1.Z}
}

// Subtract returns the difference of this vector and another vector.
func (v *Vec3) Subtract(v1 *Vec3) *Vec3 {
	return &Vec3{v.X - v1.X, v.Y - v1.Y, v.Z - v1.Z}
}

// ScalarMul returns the product with a scalar value.
func (v *Vec3) ScalarMul(s float64) *Vec3 {
	return &Vec3{v.X * s, v.Y * s, v.Z * s}
}

// ScalarDiv returns the division with a scalar value.
func (v *Vec3) ScalarDiv(s float64) *Vec3 {
	return &Vec3{v.X / s, v.Y / s, v.Z / s}
}

// Dot returns the dot product with another vector.
func (v *Vec3) Dot(v1 *Vec3) float64 {
	return v.X*v1.X + v.Y*v1.Y + v.Z*v1.Z
}

// Cross returns the cross product with another vector.
func (v *Vec3) Cross(v1 *Vec3) *Vec3 {
	return &Vec3{v.Y*v1.Z - v1.Y*v.Z, v.Z*v1.X - v1.Z*v.X, v.X*v1.Y - v1.X*v.Y}
}

// Length returns the norm of this vector.
func (v *Vec3) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Normalize returns the normalized vector.
func (v *Vec3) Normalize() *Vec3 {
	norm := v.Length()
	return &Vec3{v.X / norm, v.Y / norm, v.Z / norm}
}
