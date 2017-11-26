package main

// Intersect encapsulates all information of an intersection between a ray and an object.
type Intersect struct {
	t        float64
	position Vec3
	normal   Vec3
	material *Material
}

// NewIntersect returns a new intersect.
func NewIntersect() *Intersect {
	return &Intersect{}
}
