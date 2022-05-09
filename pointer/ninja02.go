package pointer

import "fmt"

type address struct {
	number int
	street string
	city   string
}

type person struct {
	firstname string
	lastname  string
	address   address
}

func changeMe(person *person) {
	person.address.number = 1234
}

func Ninja02() {
	person := person{
		firstname: "John",
		lastname:  "Smith",
		address: address{
			street: "Hollywood Boulevard",
			city:   "Lost Angeles",
			number: 42,
		},
	}

	fmt.Println(person)

	changeMe(&person)

	fmt.Println(person)
}
