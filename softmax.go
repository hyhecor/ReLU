package activations

import "math"

//SoftMax is exp(x) / sum(exp(x))
func SoftMax(x ...float64) []float64 {
	out := make([]float64, len(x))

	sum := func(a, b float64) float64 {
		return a + b
	}

	for i, f := range x {
		s := math.Exp(f) / Fold(sum, Map(math.Exp, x))
		out[i] = s
	}

	return out
}

//Map (higher-order function)
func Map(fn func(x float64) float64, x []float64) []float64 {
	out := make([]float64, len(x))
	for i, v := range x {
		out[i] = fn(v)
	}
	return out
}

//Fold (higher-order function)
func Fold(fn func(a, b float64) float64, x []float64) float64 {
	out := float64(0)
	for _, v := range x {
		out = fn(out, v)
	}
	return out
}

func init() {
}
