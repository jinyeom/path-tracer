package main

type Material struct {
	color *Vec3 // color

	// TODO: add more properties of material.
	//
	// For now, I'm going to stick to just color, but obviously there are more than
	// colors to a material.
}

func NewMaterial(color *Vec3) *Material {
	return &Material{
		color: color,
	}
}

// Color returns the color of the material.
func (m *Material) Color() *Vec3 {
	return m.color
}
