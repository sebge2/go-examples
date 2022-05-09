package dog

import (
	"fmt"
	"testing"
)

func TestYearsPositive(t *testing.T) {
	actual := Years(2)

	assert(t, actual, 14)
}

func TestYearsZero(t *testing.T) {
	actual := Years(0)

	assert(t, actual, 0)
}

func ExampleYears_one() {
	fmt.Println(Years(1))

	// Output:
	// 7
}

func ExampleYears_two() {
	fmt.Println(Years(2))

	// Output:
	// 14
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(i)
	}
}

func assert(t *testing.T, actual int, expected int) {
	if actual != expected {
		t.Error("expecting", expected, "got", actual)
	}
}
