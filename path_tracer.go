package main

import "math/rand"

type PathTracer struct {
	Scene  *Scene
	Camera *Camera
}

func NewPathTracer(config *Config) *PathTracer {
	pos, tangent, normal := config.Camera()
	return &PathTracer{
		Scene:  NewScene(),
		Camera: NewCamera(pos, tangent, normal),
	}
}

func (p *PathTracer) TraceAt(i, j int) Vec3 {

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
