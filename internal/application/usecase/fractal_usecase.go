package usecase

import (
	"github.com/werener/fractal-flame/internal/domain"
	"github.com/werener/fractal-flame/pkg/random"
)

type Saver interface {
	Save(fi domain.FractalImage, path string) error
}

type Generator interface {
	Generate(cfg *domain.Configuration, rnd random.RandomGenerator) domain.FractalImage
}

type FractalService struct {
	saver      Saver
	generator  Generator
	randomizer random.RandomGenerator
	image      domain.FractalImage
}

func NewFractalService(sav Saver, gen Generator, rnd random.RandomGenerator) *FractalService {
	return &FractalService{
		saver:      sav,
		generator:  gen,
		randomizer: rnd,
	}
}

func (fs *FractalService) Generate(cfg *domain.Configuration) {
	fs.image = fs.generator.Generate(cfg, fs.randomizer)
	if cfg.UseGammaCorrection {
		fs.image.GammaCorrect(cfg.Gamma)
	}
}

func (fs FractalService) Save(path string) error {
	return fs.saver.Save(fs.image, path)
}
