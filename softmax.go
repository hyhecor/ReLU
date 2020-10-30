package activations

import "math"

//SoftMax is exp(x) / sum(exp(x))
func SoftMax(x ...float64) (out []float64) {
	out = make([]float64, len(x))

	sum := func(a, b float64) float64 {
		return a + b
	}

	calc := func(a []float64) func(f float64) float64 {
		return func(e float64) float64 {
			return math.Exp(e) / Fold(float64(0), sum, Map(math.Exp, a))
		}
	}

	return Map(calc(x), x)
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
func Fold(def float64, fn func(a, b float64) float64, x []float64) float64 {
	out := def
	for _, v := range x {
		out = fn(out, v)
	}
	return out
}

func init() {
}
