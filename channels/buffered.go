package channels

import "fmt"

func BufferedExamples() {
	c := make(chan int, 1)

	go func() {
		c <- 42
		c <- 43
	}()

	fmt.Println(<-c)
	fmt.Println(<-c)
}
