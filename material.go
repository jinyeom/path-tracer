package main

type Material struct {
	color        Vec3    // color
	emissive     float64 // emissive
	ambient      float64 // ambient
	specular     float64 // specular
	diffusive    float64 // diffusive
	reflective   float64 // reflective
	transmissive float64 // transmissive
	refractIndex float64 // refractive index
	shininess    float64 // shininess
}
