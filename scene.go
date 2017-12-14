package main

// Scene is a bounded space that contains objects.
type Scene struct {
	bound   *BoundBox
	objects []Geometry
	lights  []Light
}

// NewScene creates a new Scene given its bounding box.
func NewScene(bound *BoundBox) *Scene {
	return &Scene{
		bound:   bound,
		objects: make([]Geometry, 0),
		lights:  make([]Light, 0),
	}
}

// Lights returns all of the lights in the scene.
func (s *Scene) Lights() []Light {
	return s.lights
}

// AddObject appends a new Geometry to the slice of objects.
func (s *Scene) AddObject(g Geometry) {
	s.objects = append(s.objects, g)
}

// AddLight appends a new Light to the slice of lights.
func (s *Scene) AddLight(l Light) {
	s.lights = append(s.lights, l)
}

// Intersect checks if there is any intersection between the argument ray and any object in the
// scene; return the intersection with the lowest t value, and nil if there are no intersection.
func (s *Scene) Intersect(r *Ray) *Intersect {
	minT := 10000.0 // some very large value
	isects := make(map[float64]*Intersect)
	for _, geometry := range s.objects {
		if isect := geometry.Intersect(r); isect != nil {
			t := isect.T()
			isects[t] = isect
			if t < minT {
				minT = t
			}
		}
	}
	if len(isects) == 0 {
		return nil
	}
	return isects[minT]
}
