package domain

type Configuration struct {
	Resolution     Size
	Seed           int64
	IterationCount int
	OutputPath     string
	Threads        int
	AffineParams   AffineParams
	Functions      []Function
}

type Size struct {
	Width  int
	Height int
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
	Transformation Transformation
	Weight         float64
}

var (
	WidthDefault           int       = 1920
	HeightDefault          int       = 1080
	SeedDefault            int64     = 5
	IterationCountDefault  int       = 1500
	OutputPathDefault      string    = "result.png"
	ThreadsDefault         int       = 1
	AffineParamsDefault    []float64 = []float64{0.9, 0.7, 0, -0.15, -1.1, 0}
	FunctionStringsDefault []string  = []string{"swirl:1.0"}
	GammaValueDefault      float64   = 2.2
	SymmetryLevelDefault   int       = 1
)
