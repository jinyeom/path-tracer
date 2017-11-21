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

func (b *Buffer) Width() int {
	return b.width
}

func (b *Buffer) Height() int {
	return b.height
}

// SetIntensityAt sets a pixel value at the argument offset (i, j), given its intensity of R, G,
// B, A values that range from 0.0 to 1.0.
func (b *Buffer) SetIntensityAt(i, j int, ir, ig, ib, ia float64) {
	pr := uint8(255. * math.Min(math.Max(ir, 0.0), 1.0))
	pg := uint8(255. * math.Min(math.Max(ig, 0.0), 1.0))
	pb := uint8(255. * math.Min(math.Max(ib, 0.0), 1.0))
	pa := uint8(255. * math.Min(math.Max(ia, 0.0), 1.0))

	b.buffer.SetRGBA(i, j, color.RGBA{pr, pg, pb, pa})
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
