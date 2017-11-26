package main

type Scene struct {
	bound *BoundBox
}

func NewScene() *Scene {
	return &Scene{}
}
