package random

import "math/rand"

type RandomGenerator struct {
	seed int64
}

// NewGenerator creates a new random generator with the given seed.
func NewGenerator() *RandomGenerator {
	return &RandomGenerator{}
}

// GetRandomizer returns a new randomizer instance with a set seed.
// It canbe used to generate random numbers.
func (*RandomGenerator) GetRandomizer(seed int64) Random {
	return &Randomizer{
		gen: rand.New(rand.NewSource(seed)),
	}
}

type Randomizer struct {
	gen *rand.Rand
}

func (randomizer *Randomizer) Float64() float64 {
	return randomizer.gen.Float64()
}

func (randomizer *Randomizer) Intn(n int) int {
	return randomizer.gen.Intn(n)
}
