package test

import "fmt"

func Ninja01() {
	fmt.Println("cd ./test/dog")
	fmt.Println("go test .")
	fmt.Println("go test -bench .")
	fmt.Println("go doc years")
	fmt.Println("go test -cover")
	fmt.Println("go test -coverprofile=c.out")
	fmt.Println("go tool cover -html=c.out")
}
