package activations

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSoftMax(t *testing.T) {

	sl := []float64{1.0, 2.0, 3.0, 4.0, 1.0, 2.0, 3.0}

	sum := float64(0)
	for i, softmax := range SoftMax(sl...) {
		t.Logf("SoftMax input=%f return=%f\n", sl[i], softmax)

		if sl[i] == 1.0 {
			assert.Equal(t, "2.3640543E-02", strconv.FormatFloat(softmax, 'E', -1, 32))
		}

		sum += softmax
	}
	assert.Equal(t, "1E+00", strconv.FormatFloat(sum, 'E', -1, 32))

}
