package main

type PathTracer struct {
	Scene  *Scene
	Camera *Camera

	samples int
}

func NewPathTracer(samples int) *PathTracer {
	return &PathTracer{
		Scene:   NewScene(),
		Camera:  NewCamera(),
		samples: samples,
	}
}

func (p *PathTracer) TraceAt(i, j int) Vec3 {

	return NewVec3(0.0, 0.0, 0.0)
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
