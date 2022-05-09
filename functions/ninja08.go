package functions

import "fmt"

func ninja08Provider() func(int) int {
	return func(val int) int {
		return val * val
	}
}

func Ninja08() {
	f := ninja08Provider()

	fmt.Println(f(42))
}
