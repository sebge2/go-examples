package error

import (
	"encoding/json"
	"fmt"
	"log"
)

type Ninja02Person struct {
	Firstname string
	Lastname  string
}

func Ninja02() {
	p := Ninja01Person{"John", "Smith"}

	bs, err := toJson(p)
	if err != nil {
		log.Fatalln("Cannot serialize to JSON", err)
	}

	fmt.Println(string(bs))
}

func toJson(p interface{}) ([]byte, error) {
	bs, err := json.Marshal(p)
	if err != nil {
		return []byte{}, fmt.Errorf("cannot serialize to JSON [%v]: %v", p, err)
	}

	return bs, nil
}
