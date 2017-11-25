package main

type Config struct {
	Width     int       `json:"width"`
	Height    int       `json:"height"`
	CameraPos [3]float3 `json:"cameraPos"`
}
