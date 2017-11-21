package main

// Camera casts rays to the scene through the screen to retrieve the pixel densities.
type Camera struct {
	P Vec3 // position of the camera
	D Vec3 // direction the camera is looking at
}

// NewCamera returns a new camera at (0, 0, 0) that is pointing at (0, 0, -1).
func NewCamera() *Camera {
	return &Camera{
		P: NewVec3(0.0, 0.0, 0.0),
		D: NewVec3(0.0, 0.0, -1.0),
	}
}

func (c *Camera) RayThrough(i, j int) Ray {

}
