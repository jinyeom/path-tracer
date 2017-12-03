package main

import "math/rand"

// PathTracer
type PathTracer struct {
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
	pos, tangent, normal := config.Camera()
	camera := NewCamera(pos, tangent, normal)

	return &PathTracer{
		Scene:  scene,
		Camera: camera,
	}
}

func (p *PathTracer) TraceAt(i, j int) *Vec3 {

	// TODO: implementation

	return NewVec3(rand.Float64(), rand.Float64(), rand.Float64())
}

func (p *PathTracer) Render(b *Buffer) {
	width, height := b.Dims()
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			// Set the intensity of the pixel in buffer.
			rgb := p.TraceAt(i, j)
			b.SetIntensityAt(i, j, rgb)
		}
	}
}
