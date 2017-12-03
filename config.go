package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"text/tabwriter"
	"time"
)

// Config contains the configuration from a JSON file.
type Config struct {
	// Name of the output file.
	FileName string `json:"fileName"`

	// Dimensions of the rendered image.
	Width  int `json:"width"`
	Height int `json:"height"`

	// Bounding box of the scene.
	SceneBoundMin [3]float64 `json:"sceneBoundMin"`
	SceneBoundMax [3]float64 `json:"sceneBoundMax"`

	// Camera configurations which include its position, tangent vector, and normal vector.
	// Note that its binormal vector is computed internally with the tangent and normal vectors.
	CameraPosition [3]float64 `json:"cameraPosition"`
	CameraTangent  [3]float64 `json:"cameraTangent"`
	CameraNormal   [3]float64 `json:"cameraNormal"`

	// Sample size for Monte Carlo integration. The sample size will specify how many rays with
	// random directions are sampled from a ray intersection.
	SampleSize int `json:"sampleSize"`
}

// NewDefaultConfig returns a Config with default configuration.
func NewDefaultConfig() *Config {
	return &Config{
		FileName:       fmt.Sprintf("phoebe_%d.png", time.Now().UnixNano()),
		Width:          800,
		Height:         600,
		CameraPosition: [3]float64{0.0, 0.0, 0.0},
		CameraTangent:  [3]float64{0.0, 0.0, -1.0},
		CameraNormal:   [3]float64{0.0, 1.0, 0.0},
		SampleSize:     100,
	}
}

// NewConfigJSON returns a new configuration, given an argument name of the JSON file.
func NewConfigJSON(fileName string) (*Config, error) {
	dat, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	config := Config{}
	if err = json.Unmarshal(dat, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

// Summary prints the summary of the configuration.
func (c *Config) Summary() {
	w := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', tabwriter.TabIndent)
	defer w.Flush()

	fmt.Fprintln(w, "=============== Configuration Summary ===============")
	fmt.Fprintln(w, "-----------------------------------------------------")

	// Print the output file name.
	fmt.Fprintf(w, "+ Output file name: \t%s\n", c.FileName)
	fmt.Fprintln(w, "-----------------------------------------------------")

	// Print the output image dimensions.
	fmt.Fprintf(w, "+ Image dimensions: \t(%d, %d)\n", c.Width, c.Height)
	fmt.Fprintln(w, "-----------------------------------------------------")

	// Print the camera settings.
	x, y, z := c.CameraPosition[0], c.CameraPosition[1], c.CameraPosition[2]
	tx, ty, tz := c.CameraTangent[0], c.CameraTangent[1], c.CameraTangent[2]
	nx, ny, nz := c.CameraNormal[0], c.CameraNormal[1], c.CameraNormal[2]

	fmt.Fprintln(w, "+ Camera settings:")
	fmt.Fprintf(w, "  Position: \t(%.3f, %.3f, %.3f)\n", x, y, z)
	fmt.Fprintf(w, "  Tangent: \t(%.3f, %.3f, %.3f)\n", tx, ty, tz)
	fmt.Fprintf(w, "  Normal: \t(%.3f, %.3f, %.3f)\n", nx, ny, nz)
	fmt.Fprintln(w, "-----------------------------------------------------")

	// Print the scene boundary settings.
	minX, minY, minZ := c.SceneBoundMin[0], c.SceneBoundMin[1], c.SceneBoundMin[2]
	maxX, maxY, maxZ := c.SceneBoundMax[0], c.SceneBoundMax[1], c.SceneBoundMax[2]
	fmt.Fprintln(w, "+ Scene boundary settings:")
	fmt.Fprintf(w, "  Low: \t(%.3f, %.3f, %.3f)\n", minX, minY, minZ)
	fmt.Fprintf(w, "  High: \t(%.3f, %.3f, %.3f)\n", maxX, maxY, maxZ)
	fmt.Fprintln(w, "-----------------------------------------------------")

	// Print the sample size.
	fmt.Fprintf(w, "+ Sample size: \t%d\n", c.SampleSize)
	fmt.Fprintln(w, "-----------------------------------------------------")

	fmt.Fprintln(w, "=====================================================")
}

// SceneBound returns the bounding box of the scene.
func (c *Config) SceneBound() *BoundBox {
	boundMin := NewVec3(c.SceneBoundMin[0], c.SceneBoundMin[1], c.SceneBoundMin[2])
	boundMax := NewVec3(c.SceneBoundMax[0], c.SceneBoundMax[1], c.SceneBoundMax[2])
	return NewBoundBox(boundMin, boundMax)
}

// Camera returns configurations for the camera, which are its position, its tangent vector,
// and its normal vector.
func (c *Config) Camera() (*Vec3, *Vec3, *Vec3) {
	pos := NewVec3(c.CameraPosition[0], c.CameraPosition[1], c.CameraPosition[2])
	t := NewVec3(c.CameraTangent[0], c.CameraTangent[1], c.CameraTangent[2])
	n := NewVec3(c.CameraNormal[0], c.CameraNormal[1], c.CameraNormal[2])
	return pos, t, n
}
