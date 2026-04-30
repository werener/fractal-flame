package handlers

import (
	"github.com/werener/fractal-flame/internal/domain"
	"github.com/werener/fractal-flame/pkg/random"
)

const (
	minX = -1.0
	minY = -1.0
	maxX = 1.0
	maxY = 1.0
)

type FractalGenerator struct{}

func (FractalGenerator) Generate(cfg *domain.Configuration, rnd random.RandomGenerator) domain.FractalImage {
	rect := domain.NewRectangle(minX, minY, maxX-minX, maxY-minY)
	randomizer := rnd.GetRandomizer(cfg.Seed)

	img := domain.NewFractalImage(cfg.Resolution)
	img.Generate(rect, cfg, randomizer)
	return img
}
