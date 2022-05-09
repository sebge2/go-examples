package functions

import "fmt"

func Ninja07() {

	f := func(val int) {
		fmt.Println(val)
	}

	f(42)
	fmt.Printf("%T\n", f)
}
