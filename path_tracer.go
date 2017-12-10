package main

import (
	"math/rand"
)

// PathTracer renders a 3D scene onto a 2D image.
type PathTracer struct {
	Config *Config
	Scene  *Scene
	Camera *Camera
}

// NewPathTracer constructs and returns a new path tracer according to the argument configuration.
func NewPathTracer(config *Config) *PathTracer {
	// Build the scene.
	sceneBound := config.SceneBound()
	scene := NewScene(sceneBound)

	// TODO: add objects to the scene according to the config.

	// Build the camera.
	eye, center, up := config.EyeCenterUp()
	camera := NewCamera(eye, center, up)

	return &PathTracer{
		Config: config,
		Scene:  scene,
		Camera: camera,
	}
}

// TraceAt casts a ray from the camera through a pixel at (x, y).
func (p *PathTracer) TraceAt(x, y int) *Vec3 {

	// TODO: implementation

	// Cast a ray through at pixel at (x, y).
	r := p.Camera.RayThrough(x, y, p.Config.Width, p.Config.Height)
	for _, g := range p.Scene.Objects() {
		// If there is an intersection from this ray with an object in the scene,
		if isect := g.Intersect(r); isect != nil {
			return NewVec3(rand.Float64(), rand.Float64(), rand.Float64())
		}
	}
	return NewVec3(0, 0, 0)
}

// Render traces rays through the buffer and sets each pixel value.
func (p *PathTracer) Render(b *Buffer) {
	width, height := b.Dims()
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			// Set the intensity of the pixel in buffer.
			rgb := p.TraceAt(x, y)
			b.SetIntensityAt(x, y, rgb)
		}
	}
}
