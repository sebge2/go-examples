package pointer

import "fmt"

func funcByValue(x int) {
	fmt.Println(x)

	x = 54

	fmt.Println(x)
}

func funcByPointer(x *int) {
	fmt.Println(*x)

	*x = 54

	fmt.Println(*x)
}

func PointerWhenToUseExample() {
	x := 46

	funcByValue(x)

	fmt.Println(x)
	fmt.Println()

	funcByPointer(&x)

	fmt.Println(x)
}
