package main

// Intersect encapsulates all information of an intersection between a ray and an object.
type Intersect struct {
	ray      *Ray
	geometry Geometry
	position *Vec3
	normal   *Vec3
	t        float64
}

// NewIntersect returns a new intersect.
func NewIntersect(ray *Ray, geometry Geometry, t float64) *Intersect {
	position := ray.At(t)
	normal := geometry.Normal(position)
	return &Intersect{
		ray:      ray,
		geometry: geometry,
		position: position,
		normal:   normal,
		t:        t,
	}
}

// Geometry returns the geometry on which the intersection lies.
func (i *Intersect) Geometry() Geometry {
	return i.geometry
}

// Position returns the position of the intersection.
func (i *Intersect) Position() *Vec3 {
	return i.position
}

// Normal returns the normal at the intersection.
func (i *Intersect) Normal() *Vec3 {
	return i.normal
}

// T returns the t value of the intersect.
func (i *Intersect) T() float64 {
	return i.t
}
