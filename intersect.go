package main

// Intersect encapsulates all information of an intersection between a ray and an object.
type Intersect struct {
	t        float64   // t value
	position Vec3      // position coordinates
	normal   Vec3      // normal vector
	object   Geometry  // object on which intersection happened
	material *Material // material at this intersection
}

// NewIntersect returns a new intersect.
func NewIntersect(t float64) *Intersect {
	return &Intersect{}
}
