package functions

import "fmt"

type person struct {
	firstname string
	lastname  string
}

type secretAgent struct {
	person
	ltk bool
}

func (sa secretAgent) speak() {
	fmt.Printf("Hello Iam the secret agent %v %v!\n", sa.firstname, sa.lastname)
}

func (sa person) speak() {
	fmt.Printf("Hello Iam %v %v!\n", sa.firstname, sa.lastname)
}

func MethodExample() {
	sa := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa.speak()
}
