package functions

import "fmt"

func AnonymousFunctionExample() {
	func() {
		fmt.Println("Iam in anonymous function")
	}()
}
