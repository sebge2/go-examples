package error

import "fmt"

func RecoverExamples() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovering from ", r)
		}
	}()

	fmt.Println("Calling g()")

	g(0)

	fmt.Println("Method finished, won't be displayed")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking")

		panic(fmt.Sprintf("%v", i))
	}

	defer fmt.Println("Defer in g", i)

	fmt.Println("Printing in g", i)
	g(i + 1)
}
