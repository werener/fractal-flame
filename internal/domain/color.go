package domain

import "github.com/werener/fractal-flame/pkg/random"

// Color represents a color in RGB model.
type Color struct {
	R, G, B uint32
}

func NewColor(r, g, b uint32) Color {
	return Color{R: r, G: g, B: b}
}

func RandomColor(rnd random.Random) Color {
	return NewColor(
		uint32(rnd.Intn(255)),
		uint32(rnd.Intn(255)),
		uint32(rnd.Intn(255)),
	)
}

func RandomColors(n int, rnd random.Random) []Color {
	colors := make([]Color, n)
	for i := range n {
		colors[i] = NewColor(
			uint32(rnd.Intn(255)),
			uint32(rnd.Intn(255)),
			uint32(rnd.Intn(255)),
		)
	}
	return colors
}
