package functions

import "fmt"

func ninja01Foo() int {
	return 42
}

func ninja01Bar() (int, string) {
	return 42, "42"
}

func Ninja01() {
	fmt.Println(ninja01Foo())
	fmt.Println(ninja01Bar())
}
