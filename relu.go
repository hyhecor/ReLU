package activations

import "math"

//ReLU y is always gt then 0
func ReLU(x float64) float64 {
	return math.Max(0, x)
}
