package concurrency

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func WaitGroupExample() {
	fmt.Println("Go OS \t\t", runtime.GOOS)
	fmt.Println("Go Arch \t", runtime.GOARCH)

	fmt.Println("CPU \t\t", runtime.NumCPU())
	fmt.Println("Routine \t", runtime.NumGoroutine())

	wg.Add(2)
	go counter()
	counter()

	wg.Wait()

	fmt.Println("CPU \t\t", runtime.NumCPU())
	fmt.Println("Routine \t", runtime.NumGoroutine())
}

func counter() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	wg.Done()
}
