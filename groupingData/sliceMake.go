package groupingData

import "fmt"

func SliceMakeExample() {
	x := make([]int, 5, 5)

	x[0] = 5
	x[1] = 5
	x[2] = 5

	fmt.Printf("%T\n", x)
	fmt.Printf("%v\n", x)
}
