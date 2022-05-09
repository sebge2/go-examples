package application

import (
	"encoding/json"
	"fmt"
	"os"
)

func Ninja03() {
	type person struct {
		First string
		Last  string
		Age   int
	}

	persons := []person{
		{First: "James", Last: "Bond", Age: 32},
		{First: "Miss", Last: "Moneypenny", Age: 27},
	}

	fmt.Println(persons)

	err := json.NewEncoder(os.Stdout).Encode(persons)
	if err != nil {
		fmt.Println(err)
	}
}
