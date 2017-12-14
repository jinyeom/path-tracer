package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	// Print the signature (green).
	fmt.Println("\x1B[38;5;82mPhoebe\x1B[0m Path Tracer")
	fmt.Println("Copyright (c) 2017 by Jin Yeom")

	// Import the configuration.
	//
	// If Phoebe is run with an argument configuration file name, use the config file to
	// build the path tracer. If there are errors in the process, or there is no argument
	// config file, run the path tracer with the default configuraion.
	var config *Config
	var err error
	if len(os.Args) >= 2 {
		config, err = NewConfigJSON(os.Args[1])
		if err != nil {
			fmt.Printf("\x1B[93m%s\n\x1B[0m", err)
			fmt.Println("Rendering with default configuration...")
			config = NewDefaultConfig()
		}
	} else {
		fmt.Println("Rendering with default configuration...")
		config = NewDefaultConfig()
	}

	// Print summary of the configuration.
	config.Summary()

	// Seed the random number generator.
	rand.Seed(config.Seed)

	buf := NewBuffer(config.Width, config.Height)
	tracer := NewPathTracer(config)

	// Cornell box
	// temporary!
	//
	// Creating and adding objects to the scene will be moved to configuration.
	// Adding objects should be a part of the config JSON file.
	//
	// TODO: move objects and lights to config.
	randColor1 := NewVec3(rand.Float64(), rand.Float64(), rand.Float64())
	randColor2 := NewVec3(rand.Float64(), rand.Float64(), rand.Float64())

	sphere1 := NewSphere(NewVec3(0, -6, -10), 4.0, NewMaterial(randColor1))
	sphere2 := NewSphere(NewVec3(-2, -2, -5), 2.0, NewMaterial(randColor2))
	plane1 := NewPlane(NewVec3(0, 10, -10), NewVec3(0, -1, 0), NewMaterial(NewVec3(0.7, 0.7, 0.7)))
	plane2 := NewPlane(NewVec3(0, 0, -20), NewVec3(0, 0, 1), NewMaterial(NewVec3(0.7, 0.7, 0.7)))
	plane3 := NewPlane(NewVec3(0, -10, -10), NewVec3(0, 1, 0), NewMaterial(NewVec3(0.7, 0.7, 0.7)))
	plane4 := NewPlane(NewVec3(10, 0, -10), NewVec3(-1, 0, 0), NewMaterial(NewVec3(0, 1, 0)))
	plane5 := NewPlane(NewVec3(-10, 0, -10), NewVec3(1, 0, 0), NewMaterial(NewVec3(1, 0, 0)))
	tracer.Scene.AddObject(sphere1)
	tracer.Scene.AddObject(sphere2)
	tracer.Scene.AddObject(plane1)
	tracer.Scene.AddObject(plane2)
	tracer.Scene.AddObject(plane3)
	tracer.Scene.AddObject(plane4)
	tracer.Scene.AddObject(plane5)

	// Add light sources.
	//
	// I made directional and point light, but let's just work with directional light for now...
	// It seems like directional lights are less complicated.
	//
	// TODO: test PointLight.
	light := NewDirectLight(NewVec3(1.0, 1.0, 1.0), 0.7, NewVec3(0.0, 0.0, -1.0))
	tracer.Scene.AddLight(light)

	tracer.Render(buf)
	if err := buf.ExportPNG(config.FileName); err != nil {
		fmt.Printf("\x1B[91m%s\n\x1B[0m", err)
		fmt.Println("Image export failed, exiting...")
	}
}
