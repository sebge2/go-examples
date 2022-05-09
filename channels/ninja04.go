package channels

import "fmt"

func Ninja04() {
	c := genNinja04()

	receiveNinja04(c)

	fmt.Println("about to exit")
}

func genNinja04() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}

		close(c)
	}()

	return c
}

func receiveNinja04(c <-chan int) {
	for {
		select {
		case v, ok := <-c:
			if !ok {
				fmt.Println()
				return
			}

			fmt.Printf("%v  ", v)
		}

	}
}
