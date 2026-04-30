package domain

type Configuration struct {
	Resolution      Resolution
	Seed            int64
	IterationCount  int
	PointIterations int
	OutputPath      string
	Threads         int
	AffineParams    []AffineParams
	Functions       []Function
}

type Resolution struct {
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
	WidthDefault          int       = 1920
	HeightDefault         int       = 1080
	SeedDefault           int64     = 1
	IterationCountDefault int       = 25000
	OutputPathDefault     string    = "result.png"
	ThreadsDefault        int       = 1
	AffineParamsDefault   []float64 = []float64{
		-0.6, -0.3, 0.2, -0.15, -0.7, 0.5,
		-0.3, -0.5, -0.2, 0.5, 0.7, 0.2,
		0.2, -0.35, 0.15, .25, -0.7, 0.8,
	}
	FunctionStringsDefault []string = []string{"swirl:1.0", "heart:1.0"}
	GammaValueDefault      float64  = 2.2
	SymmetryLevelDefault   int      = 1
	PointIterationsDefault int      = 70
)
