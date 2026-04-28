package png

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"testing"
)

func TestSaveValidRandomImage(t *testing.T) {
	width, height := 100, 100
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.RGBA{
				R: uint8(rand.Intn(256)),
				G: uint8(rand.Intn(256)),
				B: uint8(rand.Intn(256)),
				A: 255,
			})
		}
	}

	tempFile := "test_random.png"
	defer os.Remove(tempFile)

	err := Save(img, tempFile)
	if err != nil {
		t.Fatalf("Save() returned error: %v", err)
	}

	if _, err := os.Stat(tempFile); os.IsNotExist(err) {
		t.Fatal("Save() did not create the file")
	}

	file, err := os.Open(tempFile)
	if err != nil {
		t.Fatalf("Failed to open saved file: %v", err)
	}
	defer file.Close()

	savedImg, err := png.Decode(file)
	if err != nil {
		t.Fatalf("Failed to decode saved PNG: %v", err)
	}

	// Verify dimensions match
	if savedImg.Bounds().Dx() != width || savedImg.Bounds().Dy() != height {
		t.Errorf("Saved image dimensions mismatch: got %dx%d, want %dx%d",
			savedImg.Bounds().Dx(), savedImg.Bounds().Dy(), width, height)
	}
}

func TestSaveInvalidPath(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 10, 10))

	// Try to save to an invalid path (directory doesn't exist)
	err := Save(img, "/nonexistent/directory/test.png")
	if err == nil {
		t.Error("Save() expected error for invalid path, got nil")
	}
}

func TestSaveInvalidImage(t *testing.T) {
	path := "empty_test.png"
	img := image.NewRGBA(image.Rect(0, 0, 0, 0))

	err := Save(img, path)
	if err == nil {
		t.Error("Saving an empty image should return an error")
	}

	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Error("File wasn't deleted after write error")
		os.Remove(path)
	}
}
