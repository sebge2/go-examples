package pointer

import "fmt"

type MyTypeWithPointerReceiver struct {
	val string
}

func (t *MyTypeWithPointerReceiver) say() {
	fmt.Println(t.val)
}

func PointerMethodSetExample() {
	v := MyTypeWithPointerReceiver{
		val: "test",
	}

	v.say()
}
