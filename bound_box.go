package main

// BoundBox is a boundary defined by two points.
type BoundBox struct {
	min, max *Vec3
}

// NewBoundBox returns a new bounding box, given the two points that define the box.
func NewBoundBox(min, max *Vec3) *BoundBox {
	return &BoundBox{min, max}
}
