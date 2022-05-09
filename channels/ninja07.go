package channels

import (
	"fmt"
	"sync"
)

func Ninja07() {
	c := make(chan int)
	w := sync.WaitGroup{}

	w.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			senderNinja07(c)
			w.Done()
		}()
	}

	go func() {
		receiverNinja07(c)
	}()

	w.Wait()
	close(c)
}

func receiverNinja07(c chan int) {
	for v := range c {
		fmt.Printf("%v ", v)
	}

	fmt.Println()
}

func senderNinja07(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
}
