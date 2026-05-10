package domain

type Configuration struct {
	Resolution         Resolution     `json:"resolution"`
	Seed               int64          `json:"seed"`
	IterationCount     int            `json:"iteration_count"`
	PointIterations    int            `json:"point_iterations"`
	OutputPath         string         `json:"output_path"`
	Threads            int            `json:"threads"`
	AffineParams       []AffineParams `json:"affine_params"`
	Functions          []Function     `json:"functions"`
	UseGammaCorrection bool           `json:"gamma_correction"`
	Gamma              float64        `json:"gamma"`
	SymmetryLevel      int            `json:"symmetry_level"`
}

type Resolution struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type AffineParams struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
	C float64 `json:"c"`
	D float64 `json:"d"`
	E float64 `json:"e"`
	F float64 `json:"f"`
}

type Function struct {
	Name           TransformationType `json:"name"`
	Weight         float64            `json:"weight"`
	Transformation Transformation
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
	GammaValueDefault      float64  = 1
	SymmetryLevelDefault   int      = 1
	PointIterationsDefault int      = 70
)
