package domain

import "math"

type TransformationType string
type Transformation func(Point) Point

const (
	Swirl      TransformationType = "swirl"
	Horseshoe  TransformationType = "horseshoe"
	Spherical  TransformationType = "spherical"
	Polar      TransformationType = "polar"
	Heart      TransformationType = "heart"
	Disk       TransformationType = "disk"
	Cosine     TransformationType = "cosine"
	Sinusoidal TransformationType = "sinusoidal"
)

var AvailableTransformations = map[TransformationType]Transformation{
	Swirl:      swirl,
	Horseshoe:  horseshoe,
	Spherical:  spherical,
	Polar:      polar,
	Heart:      heart,
	Disk:       disk,
	Cosine:     cosine,
	Sinusoidal: sinusoidal}

func GetTransformation(tt TransformationType) (Transformation, bool) {
	f, ok := AvailableTransformations[tt]
	return f, ok
}

func swirl(p Point) Point {
	x, y := p.X, p.Y

	sqd := x*x + y*y
	sin := math.Sin(sqd)
	cos := math.Cos(sqd)

	return NewPoint(x*cos-y*sin, x*sin+y*cos)
}

func horseshoe(p Point) Point {
	x, y := p.X, p.Y

	d := math.Hypot(x, y)
	if d == 0 {
		return NewPoint(0, 0)
	}
	return NewPoint((x-y)*(x+y)/d, 2*x*y/d)
}

func spherical(p Point) Point {
	x, y := p.X, p.Y

	sqd := x*x + y*y
	if sqd == 0 {
		return NewPoint(0, 0)
	}
	return NewPoint(x/sqd, y/sqd)
}

func polar(p Point) Point {
	x, y := p.X, p.Y

	newX := math.Atan2(y, x) / math.Pi
	newY := math.Hypot(x, y) - 1

	return NewPoint(newX, newY)
}

func heart(p Point) Point {
	x, y := p.X, p.Y

	d := math.Hypot(x, y)
	atan := math.Atan2(y, x)
	newX := d * math.Sin(d*atan)
	newY := -d * math.Cos(d*atan)

	return NewPoint(newX, newY)
}

func disk(p Point) Point {
	x, y := p.X, p.Y

	thetaD := math.Pi * math.Hypot(x, y)
	atan := 1 / math.Pi * math.Atan2(y, x)

	return NewPoint(atan*math.Sin(thetaD), atan*math.Cos(thetaD))
}

func cosine(p Point) Point {
	x, y := p.X, p.Y

	theta := math.Pi * x
	newX := math.Cos(theta) * math.Cosh(y)
	newY := -math.Sin(theta) * math.Sinh(y)

	return NewPoint(newX, newY)
}

func sinusoidal(p Point) Point {
	x, y := p.X, p.Y

	return NewPoint(math.Sin(x), math.Sin(y))
}
