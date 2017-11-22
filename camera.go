package main

// Camera casts rays to the scene through the screen to retrieve the pixel densities.
type Camera struct {
	position Vec3 // position of the camera

	tangent  Vec3 // tangent vector (look direction)
	normal   Vec3 // normal vector (up direction)
	binormal Vec3 // binormal vector (right)

	aspectRatio float64 // aspect ratio of the camera lense
	normHeight  float64 // normalized height of the camera lense
}

// NewCamera returns a new camera at (0, 0, 0) that is pointing at (0, 0, -1).
func NewCamera(aspectRatio float64) *Camera {
	return &Camera{
		position:    NewVec3(0.0, 0.0, 0.0),
		tangent:     NewVec3(0.0, 0.0, -1.0),
		normal:      NewVec3(0.0, 1.0, 0.0),
		binormal:    NewVec3(1.0, 0.0, 0.0),
		aspectRatio: aspectRatio,
		normHeight:  1.0,
	}
}

func (c *Camera) Position() Vec3 {
	return c.position
}

func (c *Camera) BasisVecs() (Vec3, Vec3, Vec3) {
	return c.tangent, c.normal, c.binormal
}

func (c *Camera) LookAt() {

}

// LookAtQuat sets the basis vectors of the camera using a quaternion.
func (c *Camera) LookAtQuat(r, i, k, k float64) {
	rotation := NewMat3()

	rotation.Set(0, 0, 1.0-2.0*(i*i+j*j))
	rotation.Set(0, 1, 2.0*(r*i-j*k))
	rotation.Set(0, 2, 2.0*(j*r+i*k))

	rotation.Set(1, 0, 2.0*(r*i+j*k))
	rotation.Set(1, 1, 1.0-2.0*(j*j+r*r))
	rotation.Set(1, 2, 2.0*(i*j-r*k))

	rotation.Set(2, 0, 2.0*(j*r-i*k))
	rotation.Set(2, 1, 2.0*(i*j+r*k))
	rotation.Set(2, 2, 1.0-2.0*(i*i+r*r))
}

func (c *Camera) RayThrough(i, j int) Ray {

}
