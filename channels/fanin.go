package channels

import "fmt"

func FanInExample() {
	c := fanIn(boring("John"), boring("Jane"))

	for i := 0; i < 10; i++ {
		fmt.Printf(<-c)
	}

	fmt.Println()
}

func boring(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			c <- fmt.Sprintf("%s %d ", msg, i)
		}
	}()

	return c
}

func fanIn(c1, c2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			c <- <-c1
		}
	}()

	go func() {
		for {
			c <- <-c2
		}
	}()

	return c
}
