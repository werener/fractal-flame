package domain

import "math"

// Point represents a point in a continuous 2D space
type Point struct {
	X float64
	Y float64
}

func NewPoint(x, y float64) Point {
	return Point{X: x, Y: y}
}

// affineTransform applies and affine transformation
// to the point, mutating it.
func (p *Point) affineTransform(ap AffineParams) {
	p.X = p.X*ap.A + p.Y*ap.B + ap.C
	p.Y = p.X*ap.D + p.Y*ap.E + ap.F
}

// project maps the point to a pixel in the fractal image.
func (p *Point) project(rect Rectangle, fi *FractalImage) (*Pixel, bool) {
	x := int((p.X - rect.MinPoint.X) / rect.Width * float64(fi.Width))
	y := int((p.Y - rect.MinPoint.Y) / rect.Height * float64(fi.Height))

	return fi.GetPixel(x, y)
}

// rotate applies a rotation matrix to the point
func (p *Point) rotate(theta float64) {
	x := p.X
	y := p.Y

	p.X = x*math.Cos(theta) - y*math.Sin(theta)
	p.Y = x*math.Sin(theta) + y*math.Cos(theta)
}
