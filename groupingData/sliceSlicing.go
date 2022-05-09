package groupingData

import "fmt"

func SliceSlicingExample() {
	x := []int{1, 2, 3}

	for i, v := range x[2:3] {
		fmt.Println(i, v)
	}
}
