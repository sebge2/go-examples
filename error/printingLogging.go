package error

import (
	"fmt"
	"log"
	"os"
)

func PrintingLoggingExamples() {
	f, err := os.Create("/tmp/go-examples.log")
	if err != nil {
		fmt.Println("Error while creating file", err)
	}

	defer f.Close()

	log.SetOutput(f)

	fmt.Println("Check logs in ", f.Name())

	log.Println("Test logging")

	//log.Panicln("Iam panicking")
	// log.Fatalln("Iam panicking")
}
