package main

type Tracer struct {
	Scene  *Scene
	Camera *Camera
}

func NewTracer() *Tracer {
	return &Tracer{}
}

func (t *Tracer) Trace(i, j int) (float64, float64, float64, float64) {

}

func (t *Tracer) Render(b *Buffer) {
	for i := 0; i < b.Width(); i++ {
		for j := 0; j < b.Height(); j++ {
			// Set the intensity of the pixel in buffer.
			r, g, b, a := t.Trace(i, j)
			b.SetIntensityAt(i, j, r, g, b, a)
		}
	}
}
