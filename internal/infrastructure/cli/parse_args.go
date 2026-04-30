package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/urfave/cli/v3"
	"github.com/werener/fractal-flame/internal/domain"
)

// createConfig creates a configuration for the application.
// TODO: add gamma logic and config parsing
func createConfig(c *cli.Command) (*domain.Configuration, error) {
	funcs, err := parseFunctions(c.StringSlice("functions"))
	if err != nil {
		return nil, err
	}

	args := &domain.Configuration{
		Resolution: domain.Resolution{
			Width:  c.Int("width"),
			Height: c.Int("height"),
		},
		Seed:            c.Int64("seed"),
		IterationCount:  c.Int("iteration-count"),
		PointIterations: c.Int("point-iterations"),
		OutputPath:      c.String("output-path"),
		Threads:         c.Int("threads"),
		AffineParams:    parseAffine(c.Float64Slice("affine-params")),
		Functions:       funcs,
	}

	return args, nil
}

// parseAffine converts a slice of float64 values into an AffineParams struct.
func parseAffine(aps []float64) []domain.AffineParams {
	apAmount := len(aps) / 6
	affineParams := make([]domain.AffineParams, apAmount)
	for i := range apAmount {
		affineParams[i] = domain.AffineParams{
			A: aps[6*i+0],
			B: aps[6*i+1],
			C: aps[6*i+2],
			D: aps[6*i+3],
			E: aps[6*i+4],
			F: aps[6*i+5],
		}
	}
	return affineParams
}

// parseFunctions converts a slice of function names into
// a slice of transformation functions structs.
func parseFunctions(funcStrs []string) ([]domain.Function, error) {
	functions := []domain.Function{}
	for _, funcStr := range funcStrs {
		function, err := parseFunction(funcStr)
		if err != nil {
			return []domain.Function{}, err
		}

		functions = append(functions, function)
	}
	return functions, nil
}

// parseFunction converts a function name into a transformation function struct.
// It also raises an error if the function is not supported or the weight is invalid.
func parseFunction(funcStr string) (domain.Function, error) {
	transformationStr, weightStr, ok := strings.Cut(funcStr, ":")
	if !ok {
		return domain.Function{}, fmt.Errorf("%s: wrong function format - no ':'", funcStr)
	}
	transformation, ok := domain.GetTransformation(domain.TransformationType(transformationStr))
	if !ok {
		return domain.Function{}, fmt.Errorf("%s: transformation function isn't supported", transformationStr)
	}

	weight, err := strconv.ParseFloat(weightStr, 64)
	if err != nil {
		return domain.Function{}, fmt.Errorf("%s: weight must be a float number", weightStr)
	}
	if weight <= 0 {
		return domain.Function{}, fmt.Errorf("%s: weight must be positive number", weightStr)
	}

	return domain.Function{Transformation: transformation, Weight: weight}, nil
}
