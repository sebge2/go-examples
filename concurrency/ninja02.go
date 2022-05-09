package concurrency

import "fmt"

type Person struct {
	firstname string
}

func (p *Person) speak() {
	fmt.Println("Hello", p.firstname)
}

type Human interface {
	speak()
}

func saySomething(h Human) {
	h.speak()
}

func Ninja02() {
	p := Person{firstname: "John"}

	saySomething(&p)
	// doesn't work:
	// saySomething(p)

	p.speak()
}
