package channels

import (
	"fmt"
	"sync"
)

func Ninja06() {
	c := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		sendNinja06(c)
		wg.Done()
	}()

	go func() {
		receiveNinja06(c)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("")
}

func receiveNinja06(c <-chan int) {
	for v := range c {
		fmt.Printf("%v ", v)
	}
}

func sendNinja06(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}

	close(c)
}
