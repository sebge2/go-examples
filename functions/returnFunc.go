package functions

import "fmt"

func foo() string {
	return "Hello world"
}

func bar() func() string {
	return foo
}

func ReturnFuncExample() {
	fmt.Printf("Type %T\n", bar())

	fmt.Printf("Result of function %v\n", bar()())
}
