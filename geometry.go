package main

// Geometry is an interface for objects in a scene, that are visible to the camera via rays.
type Geometry interface {
	Intersect(r Ray) *Intersect
}

// Plane is defined with its position in 3D space and its normal that defines its angle.
type Plane struct {
	position *Vec3
	normal   *Vec3
	material *Material
}

// NewPlane returns a new Plane, given its position and normal.
func NewPlane(position, normal *Vec3, material *Material) *Plane {
	return &Plane{position, normal, material}
}

// Position returns the plane's position coordinates.
func (p *Plane) Position() *Vec3 {
	return p.position
}

// Normal returns the plane's normal vector.
func (p *Plane) Normal() *Vec3 {
	return p.normal
}

func (p *Plane) Material() *Material {
	return p.material
}

// Intersect checks if the argument ray intersects with the plane; if it does, returns true and
// the coordinate of intersection. Return false and an empty *Vec3 otherwise.
func (p *Plane) Intersect(r Ray) *Intersect {

	// TODO: implement ray's intersection with the plane.

	return NewIntersect(0.0)
}

type Sphere struct {
}

type Cylinder struct {
}

type Cone struct {
}
