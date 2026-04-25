package cli

import (
	"github.com/urfave/cli/v3"
	"github.com/werener/fractal-flame/internal/domain"
)

// createConfig creates a configuration of the application.
// It takes
// TODO: add gamma logic and config parsing
func createConfig(c *cli.Command) (domain.Configuration, error) {
	ap := c.Float64Slice("affine-params")
	affineParams := domain.AffineParams{
		A: ap[0],
		B: ap[1],
		C: ap[2],
		D: ap[3],
		E: ap[4],
		F: ap[5],
	}

	args := domain.Configuration{
		Size: domain.Size{
			Width:  c.Int("width"),
			Height: c.Int("height"),
		},
		Seed:           c.Int64("seed"),
		IterationCount: c.Int("iteration-count"),
		OutputPath:     c.String("output-path"),
		Threads:        c.Int("threads"),
		AffineParams:   affineParams,
		Functions:      []domain.Function{},
	}

	return args, nil
}
