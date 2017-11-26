package main

import (
	"encoding/json"
	"io/ioutil"
)

// Config contains the configuration from a JSON file.
type Config struct {
	// Dimensions of the rendered image.
	width  int `json:"width"`
	height int `json:"height"`

	// Camera configurations which include its position, tangent vector, and normal vector.
	// Note that its binormal vector is computed internally with the tangent and normal vectors.
	cameraPosition [3]float64 `json:"cameraPosition"`
	cameraTangent  [3]float64 `json:"cameraTangent"`
	cameraNormal   [3]float64 `json:"cameraNormal"`
}

// NewDefaultConfig returns a Config with default configuration.
func NewDefaultConfig() *Config {
	return &Config{
		width:          800,
		height:         600,
		cameraPosition: [3]float64{0.0, 0.0, 0.0},
		cameraTangent:  [3]float64{0.0, 0.0, -1.0},
		cameraNormal:   [3]float64{0.0, 1.0, 0.0},
	}
}

// NewConfigJSON returns a new configuration, given an argument name of the JSON file.
func NewConfigJSON(filename string) (*Config, error) {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	config := Config{}
	if err = json.Unmarshal(dat, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

// Dims returns the dimensions of the output image.
func (c *Config) Dims() (int, int) {
	return c.width, c.height
}

// CameraInfo returns configurations for the camera, which are its position, its tangent vector,
// and its normal vector.
func (c *Config) CameraInfo() (Vec3, Vec3, Vec3) {
	pos := NewVec3(c.cameraPosition[0], c.cameraPosition[1], c.cameraPosition[2])
	t := NewVec3(c.cameraTangent[0], c.cameraTangent[1], c.cameraTangent[2])
	n := NewVec3(c.cameraNormal[0], c.cameraNormal[1], c.cameraNormal[2])
	return pos, t, n
}
