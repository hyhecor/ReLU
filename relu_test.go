package activations

import "testing"

func TestRelu(t *testing.T) {

	sl := []float64{-2, -1, -0.5, -0.2, -0.1, 0.1, 0.2, 0.5, 1.0, 2.0}
	for _, f := range sl {
		relu := ReLU(f)
		t.Logf("ReLu input=%f return=%f\n", f, relu)
	}
}
