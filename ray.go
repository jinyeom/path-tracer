package main

type Ray struct {
	Position  Vec3 // position of origin of the ray
	Direction Vec3 // direction of the ray
}

// NewRay returns a new Ray with the argument position and direction.
func NewRay(position, direction Vec3) *Ray {
	return &Ray{position, direction}
}

// At returns the coordiates at the argument t value along the ray.
func (r *Ray) At(t float64) Vec3 {
	return r.Position.Add(r.Direction.ScalarMul(t))
}
