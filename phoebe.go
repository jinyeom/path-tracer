package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("\x1B[38;5;82mPhoebe\x1B[0m Path Tracer")
	fmt.Println("Copyright (c) 2017 by Jin Yeom")

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

	config.Summary()

	buf := NewBuffer(config.Width, config.Height)
	tracer := NewPathTracer(config)
	tracer.Render(buf)
	buf.ExportPNG(config.FileName)
}
