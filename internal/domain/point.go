package domain

// Point represents a point in a continuous 2D space
type Point struct {
	X float64
	Y float64
}

func NewPoint(x, y float64) Point {
	return Point{X: x, Y: y}
}
