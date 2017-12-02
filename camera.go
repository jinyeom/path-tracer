package main

// Camera casts rays to the scene through the screen to retrieve the pixel densities.
type Camera struct {
	position *Vec3 // position of the camera

	tangent  *Vec3 // tangent vector (look direction)
	normal   *Vec3 // normal vector (up direction)
	binormal *Vec3 // binormal vector (right)

	aspectRatio float64 // aspect ratio of the camera lense
	normHeight  float64 // normalized height of the camera lense
}

// NewCamera returns a new camera given its position and basis vectors. Note that only tangent and
// normal vectors are provided, since the binormal vector can be computed with tangent and normal.
func NewCamera(position, tangent, normal *Vec3) *Camera {
	return &Camera{
		position:    position,
		tangent:     tangent,
		normal:      normal,
		binormal:    tangent.Cross(normal).Normalize(),
		aspectRatio: 1.0,
		normHeight:  1.0,
	}
}

// Position returns the position coordinates of the camera.
func (c *Camera) Position() *Vec3 {
	return c.position
}

// Look returns the direction at which the camera is looking (tangent vector).
func (c *Camera) Look() *Vec3 {
	return c.tangent
}

// Up returns the up direction of the camera (normal vector).
func (c *Camera) Up() *Vec3 {
	return c.normal
}

// Right returns the right direction of the camera (binormal vector).
func (c *Camera) Right() *Vec3 {
	return c.binormal
}

func (c *Camera) LookAt() {

}

func (c *Camera) RayThrough(i, j int) *Ray {
	return NewRay(NewVec3(0, 0, 0), NewVec3(0, 0, 0))
}
