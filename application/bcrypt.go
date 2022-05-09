package application

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func BcryptExample() {
	s := "HelloWorld"

	bs, err := bcrypt.GenerateFromPassword([]byte(s), 1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))

	errComp := bcrypt.CompareHashAndPassword(bs, []byte(s))
	fmt.Println(errComp == nil)
}
