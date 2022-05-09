package controlFlow

import "fmt"

func LoopForStatement() {
	i := 1986
	for {
		if i > 2022 {
			break
		}

		fmt.Printf("%v", i)

		i++
	}
}
