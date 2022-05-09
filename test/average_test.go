package test

import (
	"fmt"
	"testing"
)

func TestAverage1and2(t *testing.T) {
	v := Average(1, 2)

	assert(t, v, 1.5)
}

func TestAverage2and2(t *testing.T) {
	v := Average(2, 2)

	assert(t, v, 1.5)
}

func ExampleAverage_positive() {
	fmt.Println(Average(4, 6))
	// Output:
	// 5
}

func ExampleAverage_negative() {
	fmt.Println(Average(-4, -6))
	// Output:
	// -5
}

func BenchmarkAverage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Average(34, 42)
	}
}

func assert(t *testing.T, v float64, expect float64) {
	if v != expect {
		t.Error("Expected", expect, "got", v)
	}
}
