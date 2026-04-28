package random

import "math/rand"

type RandomGenerator struct {
	seed int64
}

func NewGenerator(s int64) *RandomGenerator {
	return &RandomGenerator{}
}

func (*RandomGenerator) GetNumberRandomizer(seed int64) Random {
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
