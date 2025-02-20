package geometry

import "math"

type Figures int

const (
	Square Figures = iota
	Circle
	Triangle
)

func Area(f Figures) (func(float64) float64, bool) {
	switch f {
	case Square:
		return func(x float64) float64 { return x * x }, true
	case Circle:
		return func(x float64) float64 { return math.Pi * x * x }, true
	case Triangle:
		return func(x float64) float64 { return (math.Sqrt(3) / 4) * x * x }, true
	default:
		return nil, false
	}
}
