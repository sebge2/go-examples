package groupingData

import "fmt"

func SliceAppendExample() {
	x := []int{0, 1}
	x = append(x, 56, 65)

	z := []int{12, 13}
	x = append(x, z...)

	x = append(x[0:2], z[0:1]...)

	for _, v := range x {
		fmt.Println(v)
	}
}
