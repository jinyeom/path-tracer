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

func NewDiffusiveMaterial(color *Vec3) *Material {
	return &Material{
		color:        color,
		emissive:     0.0,
		ambient:      0.0,
		specular:     0.0,
		diffusive:    1.0,
		reflective:   -1.0,
		transmissive: 0.0,
		refractIndex: 0.0,
		shininess:    0.0,
	}
}

// Color returns the color of the material.
func (m *Material) Color() *Vec3 {
	return m.color
}
