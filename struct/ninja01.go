package _struct

import "fmt"

func Ninja01() {
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

	fmt.Println(p1)
	for _, v := range p1.favoriteFlavor {
		fmt.Println(v)
	}
}
