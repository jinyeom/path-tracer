package main

type Geometry interface {
	intersect(r *Ray)
}

type Sphere struct {
}

type Cylinder struct {
}

type Cone struct {
}
