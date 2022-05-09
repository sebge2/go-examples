package channels

import "fmt"

func Ninja03() {
	c := gen()

	receiveNinja03(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}

		close(c)
	}()

	return c
}

func receiveNinja03(c <-chan int) {
	for v := range c {
		fmt.Printf("%v  ", v)
	}

	fmt.Println()
}
