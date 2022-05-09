package error

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func CheckingErrorsExample() {
	writeFile()

	readFile()
}

func writeFile() {
	f, err := os.Create("/tmp/names.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	r := strings.NewReader("Wassup")

	io.Copy(f, r)
}

func readFile() {
	f, err := os.Open("/tmp/names.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}
