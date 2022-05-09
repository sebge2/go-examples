package controlFlow

import "fmt"

func ConditionalIfElseIfElse() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Printf("%v est pair\n", i)
		} else if i%3 == 0 {
			fmt.Printf("%v n'est pas pair mais divisible par 3 \n", i)
		} else {
			fmt.Printf("%v n'est pas pair\n", i)
		}
	}
}
