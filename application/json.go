package application

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
}

func JsonExample() {
	marshal()
	unmarshal()
}

func marshal() {
	person := person{
		Firstname: "John",
		Lastname:  "Doe",
	}

	v, err := json.Marshal(
		person,
	)

	if err != nil {
		fmt.Printf("Error appened %s\n", err)
	} else {
		fmt.Println(string(v))
	}
}

func unmarshal() {
	var jsonBlob = []byte(`{"firstname": "John", "lastname": "Doe"}`)
	var person person

	err := json.Unmarshal(jsonBlob, &person)
	if err != nil {
		fmt.Printf("Error appened %s\n", err)
	} else {
		fmt.Println(person)
	}
}
