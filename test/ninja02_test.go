package test

import (
	"fmt"
	"github.com/GoesToEleven/go-programming/code_samples/010-ninja-level-thirteen/02/01-code-starting/quote"
	"testing"
)

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}

func ExampleCount() {
	fmt.Println(Count("toto tata toto titi toto tata"))

	// Output:
	// 6
}

func ExampleUseCount() {
	fmt.Println(UseCount("toto tata toto titi toto tata"))

	// Output:
	// map[tata:2 titi:1 toto:3]
}

func TestCountNoWord(t *testing.T) {
	actual := Count("")

	assertInt(t, actual, 0)
}

func TestCount6WordsRepetition(t *testing.T) {
	actual := Count("toto tata toto titi toto tata")

	assertInt(t, actual, 6)
}

func TestUseCountNoWord(t *testing.T) {
	actual := UseCount("")

	assertMap(t, actual, make(map[string]int))
}

func TestUseCount6WordsRepetition(t *testing.T) {
	actual := UseCount("toto tata toto titi toto tata")

	expected := map[string]int{
		"toto": 3,
		"tata": 2,
		"titi": 1,
	}

	assertMap(t, actual, expected)
}

func assertInt(t *testing.T, actual int, expected int) {
	if actual != expected {
		t.Error("expecting", expected, "got", actual)
	}
}

func assertMap(t *testing.T, actual map[string]int, expected map[string]int) {
	if len(actual) != len(expected) {
		t.Error("expecting", expected, "got", actual)
	}

	for expectedKey, expectedValue := range expected {
		if expectedValue != actual[expectedKey] {
			t.Error("expecting", expectedValue, "got", actual[expectedKey])
		}
	}
}
