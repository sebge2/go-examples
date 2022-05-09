package test

import "testing"

func TestAverage(t *testing.T) {
	type test struct {
		data   []float64
		answer float64
	}

	tests := []test{
		test{[]float64{21, 21}, 21},
		test{[]float64{21, 21}, 42},
	}

	for _, v := range tests {
		result := Average(v.data[0], v.data[1])

		if result != v.answer {
			t.Error("Expected", v.answer, "got", result)
		}
	}
}
