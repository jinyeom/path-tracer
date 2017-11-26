package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

// Config contains the configuration from a JSON file.
type Config struct {
	// Name of the output file.
	FileName string `json:"fileName"`

	// Dimensions of the rendered image.
	Width  int `json:"width"`
	Height int `json:"height"`

	// Camera configurations which include its position, tangent vector, and normal vector.
	// Note that its binormal vector is computed internally with the tangent and normal vectors.
	CameraPosition [3]float64 `json:"cameraPosition"`
	CameraTangent  [3]float64 `json:"cameraTangent"`
	CameraNormal   [3]float64 `json:"cameraNormal"`
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

func (c *Config) Summary() {
	fmt.Println("=============== Configuration Summary ===============")

	// Print the output file name.
	fmt.Printf("+ Output file name: %s\n", c.FileName)

	// Print the output image dimensions.
	fmt.Printf("+ Image dimensions: (%d, %d)\n", c.Width, c.Height)

	// Print the camera settings.
	x, y, z := c.CameraPosition[0], c.CameraPosition[1], c.CameraPosition[2]
	tx, ty, tz := c.CameraTangent[0], c.CameraTangent[1], c.CameraTangent[2]
	nx, ny, nz := c.CameraNormal[0], c.CameraNormal[1], c.CameraNormal[2]

	fmt.Println("+ Camera settings:")
	fmt.Printf("\t* Position: \t(%.3f, %.3f, %.3f)\n", x, y, z)
	fmt.Printf("\t* Tangent: \t(%.3f, %.3f, %.3f)\n", tx, ty, tz)
	fmt.Printf("\t* Normal: \t(%.3f, %.3f, %.3f)\n", nx, ny, nz)

	fmt.Println("=====================================================")
}

// Camera returns configurations for the camera, which are its position, its tangent vector,
// and its normal vector.
func (c *Config) Camera() (Vec3, Vec3, Vec3) {
	pos := NewVec3(c.CameraPosition[0], c.CameraPosition[1], c.CameraPosition[2])
	t := NewVec3(c.CameraTangent[0], c.CameraTangent[1], c.CameraTangent[2])
	n := NewVec3(c.CameraNormal[0], c.CameraNormal[1], c.CameraNormal[2])
	return pos, t, n
}
