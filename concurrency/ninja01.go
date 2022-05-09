package concurrency

import (
	"fmt"
	"runtime"
	"sync"
)

func Ninja01() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("first")

		wg.Done()
	}()

	go func() {
		fmt.Println("second")

		wg.Done()
	}()

	fmt.Println("Number go routines", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("Number go routines", runtime.NumGoroutine())
	fmt.Println("done")
}
