package handlers

import (
	"image"
	"image/color"

	"github.com/werener/fractal-flame/internal/domain"
	"github.com/werener/fractal-flame/pkg/png"
)

func SaveFractal(fi *domain.FractalImage, path string) error {
	img := fractalToPng(fi)
	return png.Save(img, path)
}

func fractalToPng(fractal *domain.FractalImage) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, fractal.Width, fractal.Height))

	for y := range fractal.Height {
		for x := range fractal.Width {
			pixel, _ := fractal.GetPixel(x, y)
			pixelColor := color.RGBA{
				R: uint8(pixel.Color.R),
				G: uint8(pixel.Color.G),
				B: uint8(pixel.Color.B),
				A: 255,
			}

			img.Set(x, y, pixelColor)
		}
	}

	return img
}
