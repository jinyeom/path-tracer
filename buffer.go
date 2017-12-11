package main

import (
	"errors"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"path"
)

// Buffer holds intensity of each pixel in the rendered image.
type Buffer struct {
	width  int         // width of the image
	height int         // height of the image
	buffer *image.RGBA // in-memory buffer for the image
}

// NewBuffer returns a new Buffer for the rendered image.
func NewBuffer(width, height int) *Buffer {
	return &Buffer{
		width:  width,
		height: height,
		buffer: image.NewRGBA(image.Rect(0, 0, width, height)),
	}
}

// Dims returns width and height of the image buffer.
func (b *Buffer) Dims() (int, int) {
	return b.width, b.height
}

// SetIntensityAt sets the pixel value at the argument coordinates (i, j).
func (b *Buffer) SetIntensityAt(i, j int, rgb *Vec3) {
	rVal := uint8(255. * math.Min(math.Max(rgb.X, 0.0), 1.0))
	gVal := uint8(255. * math.Min(math.Max(rgb.Y, 0.0), 1.0))
	bVal := uint8(255. * math.Min(math.Max(rgb.Z, 0.0), 1.0))
	b.buffer.SetRGBA(i, j, color.RGBA{rVal, gVal, bVal, 255})
}

// ExportPNG exports the rendered image in PNG format.
func (b *Buffer) ExportPNG(filename string) error {
	if path.Ext(filename) != ".png" {
		return errors.New("phoebe: file name must have an extension of .png")
	}
	// Create a new image file with the current time as its filename.
	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		return err
	}
	// Write the image data in buffer to the file.
	err = png.Encode(f, b.buffer)
	if err != nil {
		return err
	}
	return nil
}
