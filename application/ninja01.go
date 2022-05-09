package application

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string `json:"first"`
	Age   int    `json:"age"`
}

func Ninja01() {
	u1 := user{First: "James", Age: 32}
	u2 := user{First: "Moneypenny", Age: 27}

	users := []user{u1, u2}
	fmt.Println(users)

	v, err := json.Marshal(users)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(v))
	}
}
