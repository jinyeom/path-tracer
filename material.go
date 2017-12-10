package main

type Material struct {
	color        *Vec3   // color
	emissive     float64 // emissive
	ambient      float64 // ambient
	specular     float64 // specular
	diffusive    float64 // diffusive
	reflective   float64 // reflective
	transmissive float64 // transmissive
	refractIndex float64 // refractive index
	shininess    float64 // shininess
}

// NewMaterial returns a new material.
func NewMaterial(color *Vec3) *Material {
	return &Material{
		color: color,
	}
}

// Color returns the color of the material.
func (m *Material) Color() *Vec3 {
	return m.color
}
