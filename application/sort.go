package application

import (
	"fmt"
	"sort"
)

func SortExample() {
	i := []int{42, 64, 8, 2}

	sort.Ints(i)

	fmt.Println(i)

	s := []string{"B", "A", "Test", "Z", "Yeah"}

	sort.Strings(s)

	fmt.Println(s)
}
