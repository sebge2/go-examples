package _struct

import "fmt"

func Ninja04() {
	toto := struct {
		isToto bool
	}{
		isToto: true,
	}

	fmt.Println(toto)
}
