package main

import "math"

// Light is an interface for all types of light.
type Light interface {
	// Illuminate returns the direction and intensity of the light, given the argument
	// point on which the light illuminates.
	Illuminate(point *Vec3) (*Vec3, *Vec3)
}

// light is an abstract type of the light, which contains its color and intensity.
type light struct {
	color     *Vec3
	intensity float64
}

// DirectLight represents the directional light, which inherits Light.
type DirectLight struct {
	light
	direction *Vec3
}

// NewDirectLight returns a new DirectLight, given its color, intensity, and direction.
func NewDirectLight(color *Vec3, intensity float64, direction *Vec3) *DirectLight {
	return &DirectLight{
		light:     light{color, intensity},
		direction: direction,
	}
}

// Illuminate returns the direction and intensity of the light, given the argument
// point on which the light illuminates.
func (d *DirectLight) Illuminate(point *Vec3) (*Vec3, *Vec3) {
	intensity := d.light.color.ScalarMul(d.light.intensity)
	return d.direction, intensity
}

// PointLight represents the point light, which inherits Light.
type PointLight struct {
	light
	position *Vec3
}

// NewPointLight returns a new PointLight, given its color, intensity, and position.
func NewPointLight(color *Vec3, intensity float64, position *Vec3) *PointLight {
	return &PointLight{
		light:    light{color, intensity},
		position: position,
	}
}

// Illuminate returns the direction and intensity of the light, given the argument
// point on which the light illuminates.
func (p *PointLight) Illuminate(point *Vec3) (*Vec3, *Vec3) {
	direction := point.Subtract(p.position)
	norm := direction.Length()
	direction = direction.Normalize()
	intensity := p.light.color.ScalarMul(p.light.intensity).ScalarDiv(4.0 * math.Pi * norm)
	return direction, intensity
}
