package channels

import "fmt"

func Ninja02() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
	fmt.Printf("%T\n", c)
}
