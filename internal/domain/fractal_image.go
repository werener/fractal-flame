package domain

import "github.com/werener/fractal-flame/pkg/random"

const Shift = 20 // Amount of iterations to determine the startng point

// FractalImage represents an image of a fractal.
//
// It has resolution (Width x Height) and contains a flat slice of pixels.
type FractalImage struct {
	Width  int
	Height int
	Pixels []Pixel
}

func NewFractalImage(width, height int) *FractalImage {
	return &FractalImage{
		Width: width, Height: height,
		Pixels: make([]Pixel, width*height),
	}
}

// GetPixel returns the pixel at the specified coordinates (x, y)
// if they are within the bounds of the image.
func (fi *FractalImage) GetPixel(x, y int) (*Pixel, bool) {
	if !fi.contains(x, y) {
		return nil, false
	}

	return &fi.Pixels[y*fi.Width+x], true
}

func (fi *FractalImage) contains(x, y int) bool {
	return x >= 0 && y >= 0 && x < fi.Width && y < fi.Height
}

// Generate creates the fractal image within the specified rectangle.
// TODO: make it non-reliant on configuraion.
func (fi *FractalImage) Generate(
	rect Rectangle,
	cfg *Configuration,
	color Color,
	rnd random.Random,
) {
	point := rect.RandomPoint(rnd)
	for i := range Shift + cfg.IterationCount {
		point.affineTransform(cfg.AffineParams)

		transformation := GetRandomTransformation(rnd, cfg.Functions)
		point = transformation(point)

		if i >= 0 {
			if rect.Contains(point) {
				if pixel, ok := point.project(rect, fi); ok {
					pixel.ColorPixel(color)
				}
			}
		}
	}
}
