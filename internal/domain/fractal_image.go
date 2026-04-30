package domain

const Shift = 20 // Amount of iterations to determine the startng point

// FractalImage represents a fractal image.
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
func (f *FractalImage) GetPixel(x, y int) (*Pixel, bool) {
	if !f.contains(x, y) {
		return nil, false
	}

	return &f.Pixels[y*f.Width+x], true
}

func (f *FractalImage) contains(x, y int) bool {
	return x >= 0 && y >= 0 && x < f.Width && y < f.Height
}
