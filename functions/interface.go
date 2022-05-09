package functions

type human interface {
	speak()
}

func speak(h human) {
	h.speak()
}

func InterfaceExample() {
	sa := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	p := person{
		"Miss",
		"Moneypenny",
	}

	speak(sa)
	speak(p)
}
