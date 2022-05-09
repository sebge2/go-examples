package _struct

import "fmt"

func AnonymousExample() {
	toto := struct {
		isToto bool
	}{
		isToto: true,
	}

	fmt.Println(toto)
}
