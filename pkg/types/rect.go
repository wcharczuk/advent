package types

import "fmt"

// Rect is a box defined by two effective points.
type Rect struct {
	Top    int
	Left   int
	Bottom int
	Right  int
}

// Width returns the rect width.
func (r Rect) Width() int {
	return r.Right - r.Left
}

// Height returns the rect height.
func (r Rect) Height() int {
	return r.Bottom - r.Top
}

// Area returns the rect's area.
func (r Rect) Area() int {
	return r.Height() * r.Width()
}

// Intersects returns if our rect intersects another rects.
func (r Rect) Intersects(b Rect) bool {
	xOverlap := valueInRange(r.Left, b.Left, b.Right) ||
		valueInRange(b.Left, r.Left, r.Right)
	yOverlap := valueInRange(r.Top, b.Top, b.Bottom) ||
		valueInRange(b.Top, r.Top, r.Bottom)
	return xOverlap && yOverlap
}

// String returns a string representation for the rect.
func (r Rect) String() string {
	return fmt.Sprintf("[(%d, %d) (%d, %d)]", r.Top, r.Left, r.Bottom, r.Right)
}

func valueInRange(value, min, max int) bool {
	return (value >= min) && (value < max)
}
