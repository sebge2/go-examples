package error

import (
	"encoding/json"
	"fmt"
	"log"
)

type Ninja01Person struct {
	Firstname string
	Lastname  string
}

func Ninja01() {
	p := Ninja01Person{"John", "Smith"}

	bs, err := json.Marshal(p)
	if err != nil {
		log.Fatalln("Cannot serialize to JSON", err)
	}

	fmt.Println(string(bs))
}
