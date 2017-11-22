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
	for i := 0; i < b.Width(); i++ {
		for j := 0; j < b.Height(); j++ {
			// Set the intensity of the pixel in buffer.
			rgb := p.TraceAt(i, j)
			b.SetIntensityAt(i, j, rgb)
		}
	}
}
