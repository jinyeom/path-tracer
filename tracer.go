package main

type Tracer struct {
	Scene  *Scene
	Camera *Camera
}

func NewTracer() *Tracer {
	return &Trace{}
}
