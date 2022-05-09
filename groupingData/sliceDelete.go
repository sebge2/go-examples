package groupingData

import "fmt"

func SliceDeleteExample() {
	x := []int{0, 1, 2, 3, 4, 5, 6}

	x = append(x[:2], x[3:]...)

	fmt.Println(x)
}
