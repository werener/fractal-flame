package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/werener/fractal-flame/pkg/random"
)

func TestRectangleContains(t *testing.T) {
	t.Parallel()

	r := NewRectangle(-1, -2, 1, 2)

	tests := []struct {
		name     string
		point    Point
		contains bool
	}{
		{
			name:     "in rectangle",
			point:    Point{0, 0},
			contains: true,
		},
		{
			name:     "on the vertex",
			point:    Point{-1, -2},
			contains: true,
		},
		{
			name:     "out of rectangle",
			point:    Point{2, 0},
			contains: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			contains := r.Contains(tt.point)
			require.Equal(t, tt.contains, contains)
		})
	}
}

func TestRectangleRandomPoint(t *testing.T) {
	randGen := random.NewGenerator()

	rect := NewRectangle(-1, -2, 1, 2)

	for seed := range int64(100) {
		for range 10 {
			point := rect.RandomPoint(randGen.GetRandomizer(seed))
			require.True(t, rect.Contains(point), "seed: %d. point: (%.2f, %.2f) random point should be within the rectangle", seed, point.X, point.Y)
		}
	}
}
