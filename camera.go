package main

// Camera casts rays to the scene through the screen to retrieve the pixel densities.
type Camera struct {
	position *Vec3 // position of the camera

	tangent  *Vec3 // tangent vector (look direction)
	normal   *Vec3 // normal vector (up direction)
	binormal *Vec3 // binormal vector (right direction)

	aspectRatio float64 // aspect ratio of the camera lense
	normHeight  float64 // normalized height of the camera lense
}

// NewCamera returns a new camera given its position and basis vectors. Note that only tangent and
// normal vectors are provided, since the binormal vector can be computed with tangent and normal.
func NewCamera(eye, center, up *Vec3) *Camera {
	tangent := center.Subtract(eye).Normalize()
	normal := up.Cross(tangent).Normalize()
	binormal := tangent.Cross(normal).Normalize()
	return &Camera{
		position:    eye,
		tangent:     tangent,
		normal:      normal,
		binormal:    binormal,
		aspectRatio: 1.0,
		normHeight:  1.0,
	}
}

func (c *Camera) RayThrough(i, j, width, height int) *Ray {

	return NewRay(NewVec3(0, 0, 0), NewVec3(0, 0, 0))
}
