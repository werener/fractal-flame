package cli

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var (
	errWrongDimension     error = errors.New("dimension has to be higher than zero")
	errIterationCount     error = errors.New("iteration count has to be higher than zero")
	errThreadsCount       error = errors.New("thread count has to be higher than zero")
	errAffineParams       error = errors.New("wrong amount of affine params")
	errWrongGamma         error = errors.New("gamma cannot be zero")
	errWrongSymmetryLevel error = errors.New("symmetry level has to be higher than zero")
	errWrongOutput        error = errors.New("output path has to be a writable file with .png extension")
)

func validateDimension(d int) error {
	if d < 1 {
		return errWrongDimension
	}
	return nil
}

func validateIterationCount(i int) error {
	if i < 1 {
		return errIterationCount
	}
	return nil
}

func validatePointIterations(i int) error {
	if i < 1 {
		return errIterationCount
	}
	return nil
}

func validateOutput(outPath string) error {
	ext := filepath.Ext(outPath)
	if ext != ".png" {
		return fmt.Errorf("%w: extension isn't supported: %s", errWrongOutput, ext)
	}

	dir := filepath.Dir(outPath)
	tmpFile, err := os.CreateTemp(dir, ".tmp")

	if err != nil {
		return fmt.Errorf("%w: cannot write file to directory %s: %w", errWrongOutput, dir, err)
	}
	err = tmpFile.Close()
	if err != nil {
		return err
	}
	err = os.Remove(tmpFile.Name())
	if err != nil {
		return err
	}

	return nil
}

func validateThreads(t int) error {
	if t < 1 {
		return errThreadsCount
	}
	return nil
}

func validateAffineParams(ap []float64) error {
	if len(ap)%6 != 0 {
		return fmt.Errorf("%w: %d", errAffineParams, len(ap))
	}
	return nil
}

func validateFunctions(funcStrs []string) error {
	_, err := parseFunctions(funcStrs)
	return err
}

// TODO
func validateConfig(_ string) error {
	return nil
}

func validateGamma(g float64) error {
	if g == 0 {
		return errWrongGamma
	}
	return nil
}

func validateSymmetryLevel(sl int) error {
	if sl < 1 {
		return errWrongSymmetryLevel
	}
	return nil
}
