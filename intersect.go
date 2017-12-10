package main

// Intersect encapsulates all information of an intersection between a ray and an object.
type Intersect struct {
	t        float64   // t value
	position Vec3      // position coordinates
	normal   Vec3      // normal vector
	geometry Geometry  // geometry on which intersection happened
	material *Material // material at this intersection
}

// NewIntersect returns a new intersect.
func NewIntersect(t float64, geometry Geometry) *Intersect {
	return &Intersect{
		t:        t,
		geometry: geometry,
	}
}

// T returns the t value of the intersect.
func (i *Intersect) T() float64 {
	return i.t
}
