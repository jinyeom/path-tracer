package main

import "math/rand"

type Tracer struct {
	Scene  *Scene
	Camera *Camera
}

func NewTracer() *Tracer {
	return &Tracer{}
}

func (t *Tracer) Trace(i, j int) (float64, float64, float64, float64) {
	return rand.Float64(), rand.Float64(), rand.Float64(), rand.Float64()
}

func (t *Tracer) Render(b *Buffer) {
	for i := 0; i < b.Width(); i++ {
		for j := 0; j < b.Height(); j++ {
			// Set the intensity of the pixel in buffer.
			ir, ig, ib, ia := t.Trace(i, j)
			b.SetIntensityAt(i, j, ir, ig, ib, ia)
		}
	}
}
