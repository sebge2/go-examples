package _struct

import "fmt"

func Ninja02() {
	type person struct {
		firstname      string
		lastname       string
		favoriteFlavor []string
	}

	p1 := person{
		firstname:      "James",
		lastname:       "Bond",
		favoriteFlavor: []string{"chocolate"},
	}

	m := map[string]person{
		p1.firstname: p1,
	}

	for k, v := range m {
		fmt.Printf("%v == %v\n", k, v)
	}
}
