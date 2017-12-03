package main

type Scene struct {
	bound *BoundBox
	objects []Geometry
}

func NewScene(bound *BoundBox) *Scene {
	return &Scene{
		bound: bound,
		objects: make([]Geometry, 0),
	}
}

func (s *Scene) AddObject(g Geometry) {
	s.objects = append(s.objects, g)
}
