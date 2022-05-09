package concurrency

import (
	"fmt"
	"runtime"
)

func Ninja06() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
