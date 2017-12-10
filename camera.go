package main

import (
	"math/rand"
)

// Camera casts rays to the scene through the screen to retrieve the pixel densities.
type Camera struct {
	position *Vec3 // position of the camera

	tangent  *Vec3 // tangent vector (look direction)
	normal   *Vec3 // normal vector (up direction)
	binormal *Vec3 // binormal vector (right direction)
}

// NewCamera returns a new camera given its position and basis vectors. Note that only tangent and
// normal vectors are provided, since the binormal vector can be computed with tangent and normal.
func NewCamera(eye, center, up *Vec3) *Camera {
	tangent := center.Subtract(eye).Normalize()
	normal := up.Cross(tangent).Normalize()
	binormal := tangent.Cross(normal).Normalize()
	return &Camera{
		position: eye,
		tangent:  tangent,
		normal:   normal,
		binormal: binormal,
	}
}

// RayThrough returns a ray from the camera towards a pixel position.
func (c *Camera) RayThrough(x, y, width, height int) *Ray {
	dirX := ((float64(x)+rand.Float64()-0.5)/(float64(width)-1))*2 - 1.0
	dirY := ((float64(y)+rand.Float64()-0.5)/(float64(height)-1))*2 - 1.0

	aspectRatio := float64(width) / float64(height)

	direction := c.tangent.Copy()
	direction = direction.Add(c.normal.ScalarMul(-dirX * aspectRatio))
	direction = direction.Add(c.binormal.ScalarMul(-dirY))
	direction = direction.Normalize()

	return NewRay(c.position, direction)
}
