package functions

import "fmt"

func ninja03Defer() {
	fmt.Println("After finished")
}

func Ninja03() {
	defer ninja03Defer()

	defer func() {
		fmt.Println("After finished inner")
	}()

	fmt.Println("Ninja 03")
}
