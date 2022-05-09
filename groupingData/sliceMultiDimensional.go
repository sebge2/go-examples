package groupingData

import "fmt"

func SliceMultiDimensionalExample() {
	x := [][]string{{"James", "Bond"}, {"Miss", "Moneypenny"}}

	for _, v := range (x) {
		for _, w := range (v) {
			fmt.Println(v, w)
		}
	}
}
