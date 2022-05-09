package error

import (
	"fmt"
	"log"
)

type Ninja03CustomErr struct {
	data string
}

func (c Ninja03CustomErr) Error() string {
	return fmt.Sprintf("bam %v", c.data)
}

func ninja03Foo(e error) {
	log.Println(e)
	log.Println(e.(Ninja03CustomErr).data)
}

func Ninja03() {
	err := Ninja03CustomErr{"my data"}

	ninja03Foo(err)
}
