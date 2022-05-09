package fundamentals

import "fmt"

func IotaExamples() {
	const (
		now  = 2021 + iota
		now1 = 2021 + iota
	)

	fmt.Printf("%d %d\n", now, now1)
}
