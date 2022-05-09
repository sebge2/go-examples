package functions

import "fmt"

func displaySentence(h human) {
	switch h.(type) {
	case person:
		fmt.Println("You are a person")
	case secretAgent:
		fmt.Println("You are a secret agent")
	}
}

func SwitchTypeExample() {
	sa := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	displaySentence(sa)
}
