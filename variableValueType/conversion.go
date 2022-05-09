package variableValueType

import "fmt"

func ConversionExample() {
	type myType int

	var x myType
	var y int

	x = 911
	y = int(x)

	fmt.Printf("%v %v\n", x, y)
}
