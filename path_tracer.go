package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"sync"
)

// PathTracer renders a 3D scene onto a 2D image.
type PathTracer struct {
	Config *Config
	Scene  *Scene
	Camera *Camera
}

// NewPathTracer constructs and returns a new path tracer according to the argument configuration.
func NewPathTracer(config *Config) *PathTracer {
	// Build the scene.
	sceneBound := config.SceneBound()
	scene := NewScene(sceneBound)

	// TODO: add objects to the scene according to the config.

	// Build the camera.
	eye, center, up := config.EyeCenterUp()
	camera := NewCamera(eye, center, up)

	return &PathTracer{
		Config: config,
		Scene:  scene,
		Camera: camera,
	}
}

// TraceAt casts a ray from the camera through a pixel at (x, y).
func (p *PathTracer) TraceAt(x, y int) *Vec3 {
	ray := p.Camera.RayThrough(x, y, p.Config.Width, p.Config.Height)
	return p.traceRay(ray, p.Config.TraceDepth)
}

// traceRay
func (p *PathTracer) traceRay(r *Ray, depth int) *Vec3 {
	if depth < 0 {
		// TODO: Implement environment background.
		// Not too important right now... But for sure in the future!
		return NewVec3(0.0, 0.0, 0.0)
	}
	if isect := p.Scene.Intersect(r); isect != nil {
		position := isect.Position()
		numSamples := p.Config.IntersectSampleSize

		// Total pixel intensity, initialized as black.
		indirect := NewVec3(0.0, 0.0, 0.0)

		// On the intersection point, setup a coordinate system, from which reflection rays
		// are sampled from the hemisphere on it.
		tangent, normal, binormal := createCoordSys(isect)

		// Randomly sample reflection rays.
		// Note that this process can be optimized by sampling reflections with more likely
		// angles rather than doing so completely randomly. Keep this in mind, but don't worry
		// about this right now!
		for i := 0; i < numSamples; i++ {
			a := rand.Float64()
			b := rand.Float64()
			sample := sampleHemisphere(a, b)

			// Transform the sampled ray to the world coordinates.
			worldX := sample.X*binormal.X + sample.Y*normal.X + sample.Z*tangent.X
			worldY := sample.X*binormal.Y + sample.Y*normal.Y + sample.Z*tangent.Y
			worldZ := sample.X*binormal.Z + sample.Y*normal.Z + sample.Z*tangent.Z
			sample = NewVec3(worldX, worldY, worldZ)

			indirect = indirect.Add(p.traceRay(NewRay(position, sample), depth-1).ScalarMul(a))
		}
		return indirect.ScalarDiv(float64(numSamples) / (2.0 * math.Pi))
	}
	// TODO: Implement environment background.
	// Not too important right now... But for sure in the future!
	return NewVec3(0.0, 0.0, 0.0)
}

// createCoordSys creates a coordinate system given an intersection.
// The coordinate system is returned in order of tangent, normal, and binormal.
func createCoordSys(isect *Intersect) (*Vec3, *Vec3, *Vec3) {
	var tangent, binormal *Vec3
	normal := isect.Normal()
	if math.Abs(normal.X) > math.Abs(normal.Y) {
		tangent = NewVec3(normal.Z, 0.0, -normal.X).Normalize()
	} else {
		tangent = NewVec3(0.0, -normal.Z, normal.Y).Normalize()
	}
	binormal = normal.Cross(tangent)
	return tangent, normal, binormal
}

// sampleHemisphere samples from the hemisphere distribution.
func sampleHemisphere(a, b float64) *Vec3 {
	// a = cos(theta)
	// sin^2(theta) + cos^2(theta) = 1
	// sin(theta) = sqrt(1 - cos^2(theta))
	sinTheta := math.Sqrt(1.0 - a*a)
	phi := 2.0 * math.Pi * b

	x := sinTheta * math.Cos(phi)
	z := sinTheta * math.Sin(phi)
	return NewVec3(x, a, z)
}

// Render traces rays through the buffer and sets each pixel value.
func (p *PathTracer) Render(b *Buffer) {
	numCPU := p.Config.NumCPU
	runtime.GOMAXPROCS(numCPU)

	var wg sync.WaitGroup
	wg.Add(numCPU)

	// Channel for progress bar.
	ch := make(chan int)

	// Divide the buffer into N (number of CPUs) rows.
	width, height := b.Dims()
	for i := 0; i < numCPU; i++ {
		go func(i int) {
			defer wg.Done()
			for y := i; y < height; y += numCPU {
				for x := 0; x < width; x++ {
					// Set the intensity of the pixel in buffer.
					sampleSize := p.Config.PixelSampleSize
					intensity := NewVec3(0.0, 0.0, 0.0)
					for i := 0; i < sampleSize; i++ {
						intensity = intensity.Add(p.TraceAt(x, y))
					}
					intensity = intensity.ScalarDiv(float64(sampleSize))
					b.SetIntensityAt(x, y, intensity)
				}
				// Every time a row of pixels is done, signal the progress bar.
				ch <- 1
			}
		}(i)
	}
	fmt.Println("Rendering...")
	fmt.Printf(progress(0, height))
	for i := 0; i < height; i++ {
		fmt.Printf(progress(i+<-ch, height))
	}
	fmt.Println(" ✔")
	wg.Wait()
}

// progress returns the current state of the progress bar.
func progress(curr, total int) string {
	percentage := 100 * float64(curr) / float64(total)
	str := fmt.Sprintf("\r")
	for i := 0; i < 100; i += 2 {
		if int(percentage) > i {
			str += fmt.Sprintf("\x1B[38;5;82m█\x1B[0m")
		} else if int(percentage) == i {
			str += fmt.Sprintf("\x1B[38;5;82m▆\x1B[0m")
		} else if int(percentage) == i-1 {
			str += fmt.Sprintf("\x1B[38;5;82m▃\x1B[0m")
		} else {
			str += fmt.Sprintf("\x1B[38;5;82m▁\x1B[0m")
		}
	}
	return str
}
