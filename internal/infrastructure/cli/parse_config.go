package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"github.com/urfave/cli/v3"
	"github.com/werener/fractal-flame/internal/domain"
)

func readFromJson(configPath string, command *cli.Command, cfg *domain.Configuration) error {
	var configJson domain.Configuration
	data, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, &configJson); err != nil {
		return err
	}

	// Validate affine params and functions
	for i, f := range configJson.Functions {
		transformation, ok := domain.GetTransformation(f.Name)
		if !ok {
			return fmt.Errorf("transformation function '%s' isn't supported", f.Name)
		}
		if f.Weight <= 0 {
			return fmt.Errorf("%f: weight of '%s' has to be positive number", f.Weight, f.Name)
		}
		configJson.Functions[i].Transformation = transformation
	}
	configFile := []struct {
		key       string
		value     interface{}
		validator func(interface{}) error
	}{
		{"width", configJson.Resolution.Width, func(v interface{}) error { return validateDimension(v.(int)) }},
		{"height", configJson.Resolution.Height, func(v interface{}) error { return validateDimension(v.(int)) }},
		{"seed", configJson.Seed, nil},
		{"iteration-count", configJson.IterationCount, func(v interface{}) error { return validateIterationCount(v.(int)) }},
		{"point-iterations", configJson.PointIterations, func(v interface{}) error { return validatePointIterations(v.(int)) }},
		{"output-path", configJson.OutputPath, func(v interface{}) error { return validateOutput(v.(string)) }},
		{"threads", configJson.Threads, func(v interface{}) error { return validateThreads(v.(int)) }},
		{"affine-params", configJson.AffineParams, nil},
		{"functions", configJson.Functions, nil},
		{"gamma-correction", configJson.UseGammaCorrection, nil},
		{"gamma", configJson.Gamma, func(v interface{}) error { return validateGamma(v.(float64)) }},
		{"symmetry-level", configJson.SymmetryLevel, func(v interface{}) error { return validateSymmetryLevel(v.(int)) }},
	}

	for _, field := range configFile {
		if command.IsSet(field.key) || reflect.ValueOf(field.value).IsZero() {
			continue
		}

		if field.validator != nil {
			if err := field.validator(field.value); err != nil {
				return fmt.Errorf("invalid value %v on field %s: %w", field.value, field.key, err)
			}
		}
		setFieldValue(cfg, field.key, field.value)
	}

	// fmt.Printf("CONFIG:\n%v\n", cfg)
	return nil
}

func setFieldValue(cfg *domain.Configuration, field string, value interface{}) {
	switch field {
	case "width":
		cfg.Resolution.Width = value.(int)
	case "height":
		cfg.Resolution.Height = value.(int)
	case "seed":
		cfg.Seed = value.(int64)
	case "iteration-count":
		cfg.IterationCount = value.(int)
	case "point-iterations":
		cfg.PointIterations = value.(int)
	case "output-path":
		cfg.OutputPath = value.(string)
	case "threads":
		cfg.Threads = value.(int)
	case "gamma-correction":
		cfg.UseGammaCorrection = value.(bool)
	case "gamma":
		cfg.Gamma = value.(float64)
		cfg.UseGammaCorrection = true
	case "symmetry-level":
		cfg.SymmetryLevel = value.(int)
	case "affine-params":
		cfg.AffineParams = value.([]domain.AffineParams)
	case "functions":
		cfg.Functions = value.([]domain.Function)
	default:
		fmt.Printf("Warning! No match arm for: %s\n", field)
	}
}
