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
// It can be used to generate random numbers.
func (*RandomGenerator) GetRandomizer(seed int64) Random {
	return &Randomizer{
		gen: rand.New(rand.NewSource(seed)),
	}
}

type Randomizer struct {
	gen *rand.Rand
}

// Float64 returns, as a float64, a pseudo-random number in the half-open interval [0.0,1.0).
func (randomizer *Randomizer) Float64() float64 {
	return randomizer.gen.Float64()
}

// Intn returns, as an int, a non-negative pseudo-random number in the half-open interval [0,n). It panics if n <= 0.
func (randomizer *Randomizer) Intn(n int) int {
	return randomizer.gen.Intn(n)
}
