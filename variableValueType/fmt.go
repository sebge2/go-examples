package variableValueType

import "fmt"

func FmtPackageExample() {
	x := 42
	fmt.Println(x)
	fmt.Printf("%d %b %x %T\n", x, x, x, x)
}
