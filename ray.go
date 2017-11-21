package main

type Ray struct {
	P Vec3 // point of origin of the ray
	D Vec3 // direction of the ray
}

// NewRay returns a new Ray with the argument position and direction.
func NewRay(p, d Vec3) Ray {
	return Ray{p, d}
}

func (r Ray) At(t float64) Vec3 {
	return r.P.Add(r.D.ScalarMul(t))
}
