// Package test is a bunch of functions around tests
package test

import (
	"fmt"
	"github.com/GoesToEleven/go-programming/code_samples/010-ninja-level-thirteen/02/01-code-starting/quote"
	"strings"
)

func Ninja02() {
	fmt.Println(Count(quote.SunAlso))

	for k, v := range UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}

// UseCount returns a map associating words to the number of times they appear in the specified string
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count returns the number of words in the specified string
func Count(s string) int {
	return len(strings.Fields(s))
}
