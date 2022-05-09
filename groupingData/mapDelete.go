package groupingData

import "fmt"

func MapDeleteExample() {
	x := map[string]int{
		"James Bond":      32,
		"Miss Moneypenny": 24,
	}

	delete(x, "Miss Moneypenny")

	fmt.Println(x)
}
