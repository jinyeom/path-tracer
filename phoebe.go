package main

import (
	"fmt"
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

	buf := NewBuffer(config.Width, config.Height)
	tracer := NewPathTracer(config)

	// temporary!
	sphere := NewSphere(NewVec3(0.0, 0.0, -10.0), 3.0, NewMaterial(NewVec3(1.0, 0.0, 0.0)))
	tracer.Scene.AddObject(sphere)

	tracer.Render(buf)
	if err := buf.ExportPNG(config.FileName); err != nil {
		fmt.Printf("\x1B[91m%s\n\x1B[0m", err)
		fmt.Println("Image export failed, exiting...")
	}
}
