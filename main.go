package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("\x1B[38;5;82mPhoebe\x1B[0m Path Tracer")
	fmt.Println("Copyright (c) 2017 by Jin Yeom")

	width := flag.Int("width", 800, "width of rendered image")
	height := flag.Int("height", 600, "height of rendered image")
	samples := flag.Int("samples", 5000, "number of samples")

	flag.Parse()

	tracer := NewPathTracer(*samples)
	buf := NewBuffer(*width, *height)
	tracer.Render(buf)
	buf.ExportPNG("test.png")
}
