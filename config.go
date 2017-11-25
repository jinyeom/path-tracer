package main

import (
	"encoding/json"
	"io/ioutil"
)

// Config contains the configuration from a JSON file.
type Config struct {
	// Dimensions of the rendered image.
	Width  int `json:"width"`
	Height int `json:"height"`

	// Camera configurations which include its position, tangent vector, and normal vector.
	// Note that its binormal vector is computed internally with the tangent and normal vectors.
	CameraPosition [3]float64 `json:"cameraPosition"`
	CameraTangent  [3]float64 `json:"cameraTangent"`
	CameraNormal   [3]float64 `json:"cameraNormal"`
}

// NewConfig returns a new configuration, given an argument name of the JSON file.
func NewConfig(filename string) (*Config, error) {
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
