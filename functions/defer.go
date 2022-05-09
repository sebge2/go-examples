package functions

import "fmt"

func doAtTheEnd() {
	fmt.Println("I do something at the end.")
}

func DeferExample() {
	defer doAtTheEnd()

	fmt.Println("Hello :)")
}
