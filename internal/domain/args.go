package domain

type Number float64

type Args struct {
	Size           Size
	Seed           Number
	IterationCount int
	OutputPath     string
	Threads        int
	AffineParams   AffineParams
	Functions      []Function
}

type Size struct {
	width  int
	height int
}

type AffineParams struct {
	A Number
	B Number
	C Number
	D Number
	E Number
	F Number
}

type Function struct {
	Type   Transformation
	Weight Number
}
