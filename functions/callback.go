package functions

import "fmt"

func sum(accumulator int, el int) int {
	return accumulator + el
}

func mult(accumulator int, el int) int {
	return accumulator * el
}

func reduce(accumulatorInitialValue int, f func(int, int) int, val ...int) int {
	accumulator := accumulatorInitialValue

	for _, v := range val {
		accumulator = f(accumulator, v)
	}

	return accumulator
}

func CallbackExample() {
	values := []int{1, 2, 3, 4}

	fmt.Printf("Result %v\n", reduce(0, sum, values...))
	fmt.Printf("Result %v\n", reduce(1, mult, values...))
}
