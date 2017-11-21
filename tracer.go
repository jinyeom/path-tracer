package main

type Tracer struct {
	Scene  *Scene
	Camera *Camera

	samples int
}

func NewTracer(samples int) *Tracer {
	return &Tracer{
		Scene:   NewScene(),
		Camera:  NewCamera(),
		samples: samples,
	}
}

func (t *Tracer) TraceAt(i, j int) Vec3 {

	return NewVec3(0.0, 0.0, 0.0)
}

func (t *Tracer) Render(b *Buffer) {
	for i := 0; i < b.Width(); i++ {
		for j := 0; j < b.Height(); j++ {
			// Set the intensity of the pixel in buffer.
			rgb := t.TraceAt(i, j)
			b.SetIntensityAt(i, j, rgb)
		}
	}
}
