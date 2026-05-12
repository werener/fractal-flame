package domain

import (
	"math"

	"github.com/werener/fractal-flame/pkg/random"
)

type TransformationType string
type Transformation func(*Point)

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
	Sinusoidal: sinusoidal,
}

func GetTransformation(tt TransformationType) (Transformation, bool) {
	f, ok := AvailableTransformations[tt]
	return f, ok
}

func GetRandomTransformation(rnd random.Random, functions []Function) Transformation {
	cutoff := rnd.Float64() * calculateTotalWeight(functions)
	var cumulativeWeight float64
	for _, function := range functions {
		cumulativeWeight += function.Weight
		if cumulativeWeight >= cutoff {
			return function.Transformation
		}
	}

	return nil
}

func calculateTotalWeight(functions []Function) float64 {
	var totalWeight float64
	for _, function := range functions {
		totalWeight += function.Weight
	}
	return totalWeight
}

func swirl(p *Point) {
	x, y := p.X, p.Y

	sqd := x*x + y*y
	sin := math.Sin(sqd)
	cos := math.Cos(sqd)

	p.X = x*cos - y*sin
	p.Y = x*sin + y*cos
}

func horseshoe(p *Point) {
	x, y := p.X, p.Y

	r := math.Hypot(x, y)
	if r == 0 {
		return
	}

	p.X = (x - y) * (x + y) / r
	p.Y = 2 * x * y / r
}

func spherical(p *Point) {
	sqd := p.X*p.X + p.Y*p.Y
	if sqd == 0 {
		return
	}

	p.X = p.X / sqd
	p.Y = p.Y / sqd
}

func polar(p *Point) {
	r := math.Hypot(p.X, p.Y)

	p.X = math.Atan2(p.Y, p.X) / math.Pi
	p.Y = r - 1
}

func heart(p *Point) {
	r := math.Hypot(p.X, p.Y)
	theta := math.Atan2(p.Y, p.X)

	p.X = r * math.Sin(r*theta)
	p.Y = -r * math.Cos(r*theta)
}

func disk(p *Point) {
	piR := math.Pi * math.Hypot(p.X, p.Y)
	atan := 1 / math.Pi * math.Atan2(p.Y, p.X)

	p.X = atan * math.Sin(piR)
	p.Y = atan * math.Cos(piR)
}

func cosine(p *Point) {
	piX := math.Pi * p.X

	p.X = math.Cos(piX) * math.Cosh(p.Y)
	p.Y = -math.Sin(piX) * math.Sinh(p.Y)
}

func sinusoidal(p *Point) {
	p.X = math.Sin(p.X)
	p.Y = math.Sin(p.Y)
}
