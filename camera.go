package main

// Camera
type Camera struct {
	Eye  Vec3 // position of the camera
	Look Vec3 // direction the camera is looking at
}

func NewCamera() *Camera {
	return &Camera{
		Eye:  NewVec3(0.0, 0.0, 0.0),
		Look: NewVec3(0.0, 0.0, -1.0),
	}
}

func (c *Camera) RayThrough(i, j int) Ray {

}
