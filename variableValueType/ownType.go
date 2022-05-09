package variableValueType

import "fmt"

func OwnTypeExample() {
	type myType int

	var x myType

	x = 911

	fmt.Printf("%v\n", x)
}
