package activations

import "math"

//SoftMax is exp(x) / sum(exp(x))
func SoftMax(x ...float64) (out []float64) {
	out = make([]float64, len(x))

	inputSum := SumWithBlackBox(func(f float64) float64 {
		return math.Exp(f)
	}, x...)

	for i, f := range x {
		s := math.Exp(f) / inputSum
		out[i] = s
	}
	return out
}

// SumWithBlackBox sum
func SumWithBlackBox(blackbox func(x float64) float64, x ...float64) (out float64) {
	for _, v := range x {
		out += blackbox(v)
	}
	return
}

func init() {
}
