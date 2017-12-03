package main

// Geometry is an interface for objects in a scene, that are visible to the camera via rays.
// Any thing that can intersect with a ray is a geometry.
type Geometry interface {
	Intersect(r *Ray) *Intersect
}

// Plane is defined by its position in 3D space and its normal that defines its angle.
type Plane struct {
	position *Vec3     // a point that determines the plane's position.
	normal   *Vec3     // a vector that determines the plane's angle.
	material *Material // material property of the surface
}

// NewPlane returns a new Plane, given its position and normal.
func NewPlane(position, normal *Vec3, material *Material) *Plane {
	return &Plane{position, normal, material}
}

// Intersect checks if the argument ray intersects with the plane.
// Return nil if the ray doesn't intersect.
func (p *Plane) Intersect(r *Ray) *Intersect {
	d := p.normal.Dot(r.Direction())

	// TODO: implement ray's intersection with the plane.

	return NewIntersect(0.0)
}

// Sphere is defined by its position (center), and its radius.
type Sphere struct {
	position *Vec3     // center point of the sphere
	radius   float64   // radius of the sphere
	material *Material // material property of the surface
}

// NewSphere returns a new Sphere, given its position and radius.
func NewSphere(position *Vec3, radius float64, material *Material) *Sphere {
	return &Sphere{position, radius, material}
}

// Intersect checks if the argument ray intersects with the sphere.
// Return nil if the ray doesn't intersect.
func (s *Sphere) Intersect(r *Ray) *Intersect {

	return NewIntersect(0.0)
}
