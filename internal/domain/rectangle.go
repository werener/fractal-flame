package domain

import "github.com/werener/fractal-flame/pkg/random"

// Rectangle represents a rectangle in a continuous 2D space.
// It is defined by its minimum point (minimal X and Y coordinates)
// and its width and height.
type Rectangle struct {
	MinPoint Point
	Width    float64
	Height   float64
}

func NewRectangle(x, y, width, height float64) Rectangle {
	return Rectangle{
		MinPoint: NewPoint(x, y),
		Width:    width,
		Height:   height,
	}
}

// RandomPoint returns a random point within the rectangle.
// It accepts a randomizer instance.
func (r Rectangle) RandomPoint(rnd random.Random) Point {
	x := rnd.Float64()*r.Width + r.MinPoint.X
	y := rnd.Float64()*r.Height + r.MinPoint.Y
	return NewPoint(x, y)
}
