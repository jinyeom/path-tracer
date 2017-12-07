package main

// Scene is a bounded space that contains objects.
type Scene struct {
	bound   *BoundBox
	objects []Geometry
}

// NewScene creates a new Scene given its bounding box.
func NewScene(bound *BoundBox) *Scene {
	return &Scene{
		bound:   bound,
		objects: make([]Geometry, 0),
	}
}

func (s *Scene) Bound() *BoundBox {
	return s.bound
}

func (s *Scene) Objects() []Geometry {
	return s.objects
}

// AddObject appends a new Geometry to the slice of objects.
func (s *Scene) AddObject(g Geometry) {
	s.objects = append(s.objects, g)
}
