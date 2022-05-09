package concurrency

import (
	"fmt"
	"runtime"
	"sync"
)

func Ninja04() {
	counter := 0
	mutex := sync.Mutex{}

	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			mutex.Lock()

			incrementer := counter

			runtime.Gosched()

			counter = incrementer + 1

			mutex.Unlock()

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(counter)
}
