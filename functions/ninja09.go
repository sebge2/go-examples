package functions

import "fmt"

func ninja09Do(val []int, f func(int) int) []int {
	for i, v := range val {
		val[i] = f(v)
	}

	return val
}

func Ninja09() {
	fmt.Println(
		ninja09Do(
			[]int{1, 2, 3},
			func(val int) int {
				return val * 2
			},
		),
	)
}
