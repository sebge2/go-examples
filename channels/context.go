package channels

import (
	"context"
	"fmt"
)

func ContextExample() {
	c := context.Background()

	fmt.Printf("Context: %v\n", c)
	fmt.Printf("Context Type: %T\n", c)
	fmt.Printf("Context Err: %v\n", c.Err())
	fmt.Println("-------")

	child, cancel := context.WithCancel(c)

	fmt.Printf("Context: %v\n", child)
	fmt.Printf("Context Type: %T\n", child)
	fmt.Printf("Context Err: %v\n", child.Err())
	fmt.Println("-------")

	cancel()

	fmt.Printf("Context: %v\n", child)
	fmt.Printf("Context Type: %T\n", child)
	fmt.Printf("Context Err: %v\n", child.Err())
	fmt.Println("-------")
}
