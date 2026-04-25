package cli

import (
	"context"

	"github.com/urfave/cli/v3"
	"github.com/werener/fractal-flame/internal/domain"
)

func Run(ctx context.Context, args []string) error {
	app := &cli.Command{
		Name:     "fractal-flame",
		Usage:    "Generates fractal flames",
		HideHelp: true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "help",
				Usage: "Show help",
			},
			&cli.IntFlag{
				Name:      "width",
				Aliases:   []string{"w"},
				Value:     domain.WidthDefault,
				Usage:     "Width of the produced image",
				Validator: validateDimension,
			},
			&cli.IntFlag{
				Name:      "height",
				Aliases:   []string{"h"},
				Value:     domain.HeightDefault,
				Usage:     "Height of the produced image",
				Validator: validateDimension,
			},
			&cli.Int64Flag{
				Name:  "seed",
				Value: domain.SeedDefault,
				Usage: "Seed value for the random generation",
			},
			&cli.IntFlag{
				Name:      "iteration-count",
				Aliases:   []string{"i"},
				Value:     domain.IterationCountDefault,
				Usage:     "Number of iterations during image generation",
				Validator: validateIterationCount,
			},
			&cli.StringFlag{
				Name:      "output-path",
				Aliases:   []string{"o"},
				Value:     domain.OutputPathDefault,
				Usage:     "Path to the PNG output file",
				Validator: validateOutput,
			},
			&cli.IntFlag{
				Name:      "threads",
				Aliases:   []string{"t"},
				Value:     domain.ThreadsDefault,
				Usage:     "Number of threads to use during image generation",
				Validator: validateThreads,
			},
			&cli.Float64SliceFlag{
				Name:      "affine-params",
				Aliases:   []string{"ap"},
				Value:     domain.AffineParamsDefault,
				Usage:     "Parameters of the affine transform <a1>,<b1>,<c1>,<d1>,<e1>,<f1>/.../<aN>,<bN>,<cN>,<dN>,<eN>,<fN>",
				Validator: validateAffineParams,
			},
			&cli.StringSliceFlag{
				Name:      "functions",
				Aliases:   []string{"f"},
				Value:     domain.FunctionStringsDefault,
				Usage:     "Transform functions: <func>:<weight>,<func>:<weight>,...",
				Validator: validateFunctions,
			},
			&cli.StringFlag{
				Name:      "config",
				Usage:     "Path to json config file",
				Validator: validateConfig,
			},
			&cli.BoolFlag{
				Name:    "gamma-correction",
				Aliases: []string{"g"},
				Usage:   "Enables gamma correction",
			},
			&cli.Float64Flag{
				Name:      "gamma",
				Value:     domain.GammaValueDefault,
				Usage:     "Gamma value for brightness correction of final image",
				Validator: validateGamma,
			},
			&cli.IntFlag{
				Name:      "symmetry-level",
				Aliases:   []string{"s"},
				Value:     domain.SymmetryLevelDefault,
				Usage:     "Amount of symmetry parts in final image",
				Validator: validateSymmetryLevel,
			},
		},
		Action: executeCommand,
	}

	if err := app.Run(ctx, args); err != nil {

	}
	return nil
}
func executeCommand(_ context.Context, c *cli.Command) error {
	return nil
}
