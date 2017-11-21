package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("\x1B[38;5;82mPhoebe\x1B[0m Path Tracer")
	fmt.Println("Copyright (c) 2017 by Jin Yeom")

	width := flag.Int("width", 800, "image width")
	height := flag.Int("height", 600, "image height")
	flag.Parse()

	// tracer := NewTracer()
}
