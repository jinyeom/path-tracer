package main

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
		return NewVec3(0.0, 0.0, 0.0)
	}
	if isect := p.Scene.Intersect(r); isect != nil {
		position := isect.Position()

		// total pixel intensity
		intensity := NewVec3(0.0, 0.0, 0.0)

		// Randomly sample reflection rays.
		// Note that this process can be optimized by sampling reflections with more likely
		// angles rather than doing so completely randomly. Keep this in mind, but don't worry
		// about this right now!
		for i := 0; i < p.Config.IntersectSampleSize; i++ {
			// reflective ray

			// refractive ray
		}
		return isect.Geometry().Material().Color()
	}
	// TODO: Implement environment background.
	// Not too important right now... But for sure in the future!
	return NewVec3(0.0, 0.0, 0.0)
}

// Render traces rays through the buffer and sets each pixel value.
func (p *PathTracer) Render(b *Buffer) {
	width, height := b.Dims()
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			// Set the intensity of the pixel in buffer.
			sampleSize := p.Config.PixelSampleSize
			intensity := NewVec3(0.0, 0.0, 0.0)
			for i := 0; i < sampleSize; i++ {
				intensity = intensity.Add(p.TraceAt(x, y))
			}
			intensity = intensity.ScalarDiv(float64(sampleSize))
			b.SetIntensityAt(x, y, intensity)
		}
	}
}
