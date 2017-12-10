package main

import (
	"fmt"
	"math"
)

// Geometry is an interface for objects in a scene, that are visible to the camera via rays.
// Any thing that can intersect with a ray is a geometry.
type Geometry interface {
	String() string
	Material() *Material
	Intersect(r *Ray) *Intersect
}

// Triangle is a Geometry that is defined by 3 points and its normal.
type Triangle struct {
	a, b, c  *Vec3     // three points that define the triangle
	normal   *Vec3     // normal vector of the triangle
	material *Material // material property of the triangle surface
}

// NewTriangle returns a new Triangle given three points that define the triangle.
func NewTriangle(a, b, c, normal *Vec3, material *Material) *Triangle {
	return &Triangle{a, b, c, normal, material}
}

// String returns the string representation of the triangle.
func (t *Triangle) String() string {
	a := fmt.Sprintf("a=(%.3f, %.3f, %.3f)", t.a.X, t.a.Y, t.a.Z)
	b := fmt.Sprintf("b=(%.3f, %.3f, %.3f)", t.b.X, t.b.Y, t.b.Z)
	c := fmt.Sprintf("c=(%.3f, %.3f, %.3f)", t.c.X, t.c.Y, t.c.Z)
	return fmt.Sprintf("Triangle(%s, %s, %s)", a, b, c)
}

// Material returns the triangle's material.
func (t *Triangle) Material() *Material {
	return t.material
}

func (t *Triangle) Intersect(r *Ray) *Intersect {
	// edgeAB := t.b.Subtract(t.a)
	// edgeAC := t.c.Subtract(t.a)

	return NewIntersect(0.0, nil)
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

func (p *Plane) String() string {
	return fmt.Sprintf("Plane(p=(%.3f, %.3f, %.3f), n=(%.3f, %.3f, %.3f))",
		p.position.X, p.position.Y, p.position.Z, p.normal.X, p.normal.Y, p.normal.X)
}

// Material returns the plane's material.
func (p *Plane) Material() *Material {
	return p.material
}

// Intersect checks if the argument ray intersects with the plane.
// Return nil if the ray doesn't intersect.
func (p *Plane) Intersect(r *Ray) *Intersect {
	d := p.normal.Dot(r.Direction())
	if math.Abs(d) < epsilon {
		return nil
	}
	dist := p.position.Subtract(r.Position())
	t := dist.Dot(p.normal) / d
	if t < epsilon {
		return nil
	}
	return NewIntersect(t, p)
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

// String returns the string representation of the sphere.
func (s *Sphere) String() string {
	return fmt.Sprintf("Sphere(p=(%.3f, %.3f, %.3f), r=%.3f)",
		s.position.X, s.position.Y, s.position.Z, s.radius)
}

// Material returns the sphere's material.
func (s *Sphere) Material() *Material {
	return s.material
}

// Intersect checks if the argument ray intersects with the sphere.
// Return nil if the ray doesn't intersect.
func (s *Sphere) Intersect(r *Ray) *Intersect {
	v := r.Position().Subtract(s.position)
	b := v.Dot(r.Direction())
	c := v.Dot(v) - s.radius*s.radius
	disc := b*b - c
	if disc > 0.0 {
		disc = math.Sqrt(disc)
		if t1 := -b + disc; t1 > epsilon {
			return NewIntersect(t1, s)
		}
		if t2 := -b - disc; t2 > epsilon {
			return NewIntersect(t2, s)
		}
	}
	return nil
}
