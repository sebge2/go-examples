package test

import (
	"fmt"
	"github.com/GoesToEleven/go-programming/code_samples/010-ninja-level-thirteen/03/01-code-starting/mymath"
	"testing"
)

func BenchmarkCenteredAvg(t *testing.B) {
	for i := 0; i < t.N; i++ {
		mymath.CenteredAvg([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10000})
	}
}

func ExampleCenteredAvg() {
	fmt.Println(mymath.CenteredAvg([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10000}))

	// Output:
	// 5.5
}

func TestCenteredAvg10Numbers(t *testing.T) {
	actual := mymath.CenteredAvg([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10000})

	assertFloat(t, actual, 4)
}

func TestCenteredAvg4Numbers(t *testing.T) {
	actual := mymath.CenteredAvg([]int{1, 2, 4, 6, 9})

	assertFloat(t, actual, 5.5)
}

func assertFloat(t *testing.T, actual float64, expected float64) {
	if actual != expected {
		t.Error("expecting", expected, "got", actual)
	}
}
