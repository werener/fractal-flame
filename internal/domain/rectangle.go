package domain

type Rectangle struct {
	MinPoint Point
	Width    float64
	Height   float64
}

func NewRectangle(x, y, width, height float64) Rectangle {
	return Rectangle{
		MinPoint: NewPoint(x, y),
		Width:    width,
		Height:   height,
	}
}
