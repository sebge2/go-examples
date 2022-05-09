package channels

import "fmt"

func DirectionalExamples() {
	bidirectional := make(chan int)
	sendOnly := make(chan<- int, 1)
	receiveOnly := make(<-chan int)

	fmt.Printf("%T\n", bidirectional)
	fmt.Printf("%T\n", sendOnly)
	fmt.Printf("%T\n", receiveOnly)

	go func() {
		// cannot do:
		// receiveOnly <- 42

		sendOnly <- 42
		sendOnly <- 43
	}()

	// cannot do:
	// fmt.Println(<-sendOnly)

	// allowed:
	// fmt.Println(<-receiveOnly)
}
