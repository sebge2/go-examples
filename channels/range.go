package channels

import "fmt"

func RangeExample() {
	c := make(chan int)

	go sender(c)

	for i := range c {
		fmt.Printf("%v ", i)
	}
}

func sender(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}

	close(c)
}
