package png

import (
	"image"
	"image/png"
	"os"
)

// Save takes in an image and a path.
// It creates a file at path and writes the image to it.
func Save(img image.Image, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		os.Remove(path)
		return err
	}
	return nil
}
