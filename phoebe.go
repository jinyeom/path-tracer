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
			fmt.Println(err)
			fmt.Println("Rendering with default configuration...")
			config = NewDefaultConfig()
		}
	} else {
		fmt.Println("Rendering with default configuration...")
		config = NewDefaultConfig()
	}

	// Print summary of the configuration.
	config.Summary()

	buf := NewBuffer(config.Width, config.Height) // Create an image buffer.
	tracer := NewPathTracer(config)               // Build a path tracer with the configuration.
	tracer.Render(buf)                            // With the path tracer, render onto the buffer.
	buf.ExportPNG(config.FileName)                // Export the image as a PNG file.
}
