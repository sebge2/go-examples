package application

import (
	"fmt"
	"sort"
)

type personNinja05 struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []personNinja05

func (p ByAge) Len() int {
	return len(p)
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func Ninja05() {
	persons := []personNinja05{
		{First: "James", Last: "Bond", Age: 32, Sayings: []string{"Z", "a"}},
		{First: "Miss", Last: "Moneypenny", Age: 27, Sayings: []string{"b", "a"}},
	}

	sort.Sort(ByAge(persons))

	for _, person := range persons {
		fmt.Println(person.First, person.Last)

		sort.Strings(person.Sayings)
		for _, saying := range person.Sayings {
			fmt.Println("\t", saying)
		}
	}
}
