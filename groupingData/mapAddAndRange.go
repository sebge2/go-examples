package groupingData

import "fmt"

func MapAddAndRangeExample() {
	x := map[string]int{
		"James Bond":      32,
		"Miss Moneypenny": 24,
	}

	x["Seb"] = 36

	for p, a := range x {
		fmt.Printf("%v is %v\n", p, a)
	}
}
