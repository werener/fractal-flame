package domain

// Color represents a color in RGB model.
type Color struct {
	R, G, B uint32
}

func NewColor(r, g, b uint32) Color {
	return Color{R: r, G: g, B: b}
}
