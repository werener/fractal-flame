package domain

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSwirl(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    Point
		expected Point
	}{
		{
			name:     "origin point",
			input:    NewPoint(0, 0),
			expected: NewPoint(0, 0),
		},
		{
			name:     "point (1, 0)",
			input:    NewPoint(1, 0),
			expected: NewPoint(math.Cos(1), math.Sin(1)),
		},
		{
			name:     "point (0, 1)",
			input:    NewPoint(0, 1),
			expected: NewPoint(-math.Sin(1), math.Cos(1)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			swirl(&tt.input)
			require.InDelta(t, tt.expected.X, tt.input.X, 1e-9)
			require.InDelta(t, tt.expected.Y, tt.input.Y, 1e-9)
		})
	}
}

func TestHorseshoe(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    Point
		expected Point
	}{
		{
			name:     "origin point",
			input:    NewPoint(0, 0),
			expected: NewPoint(0, 0),
		},
		{
			name:     "point (1, 0)",
			input:    NewPoint(1, 0),
			expected: NewPoint(1, 0),
		},
		{
			name:     "point (0, 1)",
			input:    NewPoint(0, 1),
			expected: NewPoint(-1, 0),
		},
		{
			name:     "point (1, 1)",
			input:    NewPoint(1, 1),
			expected: NewPoint(0, 2/math.Sqrt2),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			horseshoe(&tt.input)
			require.InDelta(t, tt.expected.X, tt.input.X, 1e-9)
			require.InDelta(t, tt.expected.Y, tt.input.Y, 1e-9)
		})
	}
}

func TestSpherical(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    Point
		expected Point
	}{
		{
			name:     "origin point",
			input:    NewPoint(0, 0),
			expected: NewPoint(0, 0),
		},
		{
			name:     "point (1, 0)",
			input:    NewPoint(1, 0),
			expected: NewPoint(1, 0),
		},
		{
			name:     "point (0, 1)",
			input:    NewPoint(0, 1),
			expected: NewPoint(0, 1),
		},
		{
			name:     "point (2, 0)",
			input:    NewPoint(2, 0),
			expected: NewPoint(0.5, 0),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			spherical(&tt.input)
			require.InDelta(t, tt.expected.X, tt.input.X, 1e-9)
			require.InDelta(t, tt.expected.Y, tt.input.Y, 1e-9)
		})
	}
}

func TestSinusoidal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    Point
		expected Point
	}{
		{
			name:     "origin point",
			input:    NewPoint(0, 0),
			expected: NewPoint(0, 0),
		},
		{
			name:     "point (pi/2, 0)",
			input:    NewPoint(math.Pi/2, 0),
			expected: NewPoint(1, 0),
		},
		{
			name:     "point (0, pi/2)",
			input:    NewPoint(0, math.Pi/2),
			expected: NewPoint(0, 1),
		},
		{
			name:     "point (pi, pi)",
			input:    NewPoint(math.Pi, math.Pi),
			expected: NewPoint(0, 0),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sinusoidal(&tt.input)
			require.InDelta(t, tt.expected.X, tt.input.X, 1e-9)
			require.InDelta(t, tt.expected.Y, tt.input.Y, 1e-9)
		})
	}
}

func TestHeart(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    Point
		expected Point
	}{
		{
			name:     "origin point",
			input:    NewPoint(0, 0),
			expected: NewPoint(0, -0),
		},
		{
			name:     "point (1, 0)",
			input:    NewPoint(1, 0),
			expected: NewPoint(1, 0),
		},
		{
			name:     "point (0, 1)",
			input:    NewPoint(0, 1),
			expected: NewPoint(0, -1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heart(&tt.input)
			require.InDelta(t, tt.expected.X, tt.input.X, 1e-9)
			require.InDelta(t, tt.expected.Y, tt.input.Y, 1e-9)
		})
	}
}

func TestCosine(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    Point
		expected Point
	}{
		{
			name:     "origin point",
			input:    NewPoint(0, 0),
			expected: NewPoint(1, 0),
		},
		{
			name:     "point (1, 0)",
			input:    NewPoint(1, 0),
			expected: NewPoint(-1, 0),
		},
		{
			name:     "point (0, 1)",
			input:    NewPoint(0, 1),
			expected: NewPoint(math.Cosh(1), 0),
		},
		{
			name:     "point (0.5, 0.5)",
			input:    NewPoint(0.5, 0.5),
			expected: NewPoint(0, -math.Sinh(0.5)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cosine(&tt.input)
			require.InDelta(t, tt.expected.X, tt.input.X, 1e-9)
			require.InDelta(t, tt.expected.Y, tt.input.Y, 1e-9)
		})
	}
}
