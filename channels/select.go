package channels

import "fmt"

func SelectExample() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(eve, odd, quit)

	receive(eve, odd, quit)
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}

	q <- 0
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-o:
			fmt.Printf("odd %d ", v)
		case v := <-e:
			fmt.Printf("eve %d ", v)
		case v, ok := <-q:
			if ok {
				fmt.Printf("quit %d\n", v)
			} else {
				fmt.Printf("quit KO %d\n", v)
			}

			return
		}
	}
}
