package concurrency

import (
	"fmt"
	"runtime"
	"sync"
)

func Ninja03() {
	counter := 0

	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			incrementer := counter

			runtime.Gosched()

			counter = incrementer + 1

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(counter)
}
