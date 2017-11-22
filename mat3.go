package main

// Mat3 is a 3x3 matrix that contains its data in an array with the size of 9.
type Mat3 struct {
	data [9]float64
}

// NewZeros returns a 3x3 zero matrix.
func NewZeros() Mat3 {
	return Mat3{
		data: [9]float64{
			0.0, 0.0, 0.0,
			0.0, 0.0, 0.0,
			0.0, 0.0, 0.0},
	}
}

// NewIdentity returns a 3x3 identity matrix.
func NewIdentity() Mat3 {
	return Mat3{
		data: [9]float64{
			1.0, 0.0, 0.0,
			0.0, 1.0, 0.0,
			0.0, 0.0, 1.0},
	}
}

// At returns the value at (i, j).
func (m Mat3) At(i, j int) float64 {
	return m.data[i*3+j]
}

// Set sets the argument value at (i, j).
func (m Mat3) Set(i, j int, v float64) {
	m.data[i*3+j] = v
}

// SetRow sets the argument Vec3 as a row at i.
func (m Mat3) SetRow(i int, v Vec3) {
	m.Set(i, 0, v.X)
	m.Set(i, 1, v.Y)
	m.Set(i, 2, v.Z)
}

// SetCol sets the argument Vec3 as a column at j.
func (m Mat3) SetCol(j int, v Vec3) {
	m.Set(0, j, v.X)
	m.Set(1, j, v.Y)
	m.Set(2, j, v.Z)
}

// DotVec3 returns a dot product with the argument vector.
func (m Mat3) DotVec3(v Vec3) Vec3 {

}
