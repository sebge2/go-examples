package application

import (
	"encoding/json"
	"io"
	"os"
)

func WriterInterfaceExample() {
	io.WriteString(os.Stdout, "Hello\n")

	person := person{
		Firstname: "John",
		Lastname:  "Doe",
	}

	json.NewEncoder(os.Stdout).Encode(person)
}
