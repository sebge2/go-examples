package application

import (
	"fmt"
	"sort"
)

type ByFirstName []person

func (a ByFirstName) Len() int {
	return len(a)
}

func (a ByFirstName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByFirstName) Less(i, j int) bool {
	return a[i].Firstname < a[j].Firstname
}

func SortCustomExample() {

	persons := []person{
		{Firstname: "Robert", Lastname: "Peters"},
		{Firstname: "John", Lastname: "Smith"},
		{Firstname: "Julia", Lastname: "Smith"},
	}

	sort.Sort(ByFirstName(persons))

	fmt.Println(persons)
}
