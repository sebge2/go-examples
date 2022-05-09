package functions

import "fmt"

type Ninja04Person struct {
	firstname string
	lastname  string
	age       int
}

func (p Ninja04Person) speak() {
	fmt.Printf("Iam %v %v and Iam %v old\n", p.firstname, p.lastname, p.age)
}

func Ninja04() {
	p := Ninja04Person{
		lastname:  "Bond",
		firstname: "James",
		age:       32,
	}

	p.speak()
}
