package cli

import (
	"errors"
	"fmt"
)

var (
	errWrongDimension     error = errors.New("dimension has to be higher than zero")
	errIterationCount     error = errors.New("iteration count has to be higher than zero")
	errThreadsCount       error = errors.New("thread count has to be higher than zero")
	errAffineParams       error = errors.New("wrong amount of affine params")
	errWrongGamma         error = errors.New("gamma cannot be zero")
	errWrongSymmetryLevel error = errors.New("symmetry level has to be higher than zero")
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

// TODO
func validateOutput(_ string) error {
	return nil
}

func validateThreads(t int) error {
	if t < 1 {
		return errThreadsCount
	}
	return nil
}

func validateAffineParams(ap []float64) error {
	if len(ap) != 6 {
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
