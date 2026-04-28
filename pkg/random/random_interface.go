package random

type Random interface {
	Float64() float64
	Intn(n int) int
}
