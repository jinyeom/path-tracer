package main

type Material struct {
	color    *Vec3   // color
	emissive float64 // emissiveness
}

func NewMaterial(color *Vec3, emissive float64) *Material {
	return &Material{
		color:    color,
		emissive: emissive,
	}
}

func NewDiffusiveMaterial(color *Vec3) *Material {
	return &Material{
		color:    color,
		emissive: 0.0,
	}
}

// Color returns the color of the material.
func (m *Material) Color() *Vec3 {
	return m.color
}

func (m *Material) Emissive() float64 {
	return m.emissive
}
