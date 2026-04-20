package domain

type Args struct {
	Size           Size
	seed           float64
	iterationCount int
	outputPath     string
	threads        int
	affineParams   AffineParams
	Functions      []Function
}

type Size struct {
	width  int
	height int
}

type AffineParams struct {
	A float64
	B float64
	C float64
	D float64
	E float64
	F float64
}

type Function struct {
	Type   Transformation
	Weight float64
}
