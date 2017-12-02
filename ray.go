package main

type Ray struct {
	position  *Vec3 // position of origin of the ray
	direction *Vec3 // direction of the ray
}

// NewRay returns a new Ray with the argument position and direction.
func NewRay(position, direction *Vec3) *Ray {
	return &Ray{position, direction}
}

// Position returns the position of the ray.
func (r *Ray) Position() *Vec3 {
	return r.position
}

// Direction returns the direction of the ray.
func (r *Ray) Direction() *Vec3 {
	return r.direction
}

// At returns the coordiates at the argument t value along the ray.
func (r *Ray) At(t float64) *Vec3 {
	return r.position.Add(r.direction.ScalarMul(t))
}
