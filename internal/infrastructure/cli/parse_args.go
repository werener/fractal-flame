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
		Size: domain.Size{
			Width:  c.Int("width"),
			Height: c.Int("height"),
		},
		Seed:           c.Int64("seed"),
		IterationCount: c.Int("iteration-count"),
		OutputPath:     c.String("output-path"),
		Threads:        c.Int("threads"),
		AffineParams:   parseAffine(c.Float64Slice("affine-params")),
		Functions:      funcs,
	}

	return args, nil
}

func parseAffine(aps []float64) domain.AffineParams {
	return domain.AffineParams{
		A: aps[0],
		B: aps[1],
		C: aps[2],
		D: aps[3],
		E: aps[4],
		F: aps[5],
	}
}

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

	return domain.Function{Function: transformation, Weight: weight}, nil
}
