package functions

import "fmt"

func ninja02Foo(val ...int) int {
	sum := 0

	for _, el := range val {
		sum += el
	}

	return sum
}

func ninja02Bar(val []int) int {
	sum := 0

	for _, el := range val {
		sum += el
	}

	return sum
}

func Ninja02() {
	x := []int{1, 2, 3}

	fmt.Println(ninja02Foo(x...))
	fmt.Println(ninja02Bar(x))
}
