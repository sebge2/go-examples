package functions

import "fmt"

func FuncExpressionExample() {
	f := func(val int) {
		fmt.Printf("Iam in a func expression %v\n", val)
	}

	f(42)
}
