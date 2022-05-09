package groupingData

import "fmt"

func SliceForRangeExample() {
	x := []int{1, 2, 3}

	for _, v := range x {
		fmt.Println(v)
	}
}
