package _struct

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle

	fourWheel bool
}

type sedan struct {
	vehicle

	luxury bool
}

func EmbeddedExample() {
	t := truck{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "Yellow",
		},
		luxury: true,
	}

	fmt.Println(t)
	fmt.Println(s)
}
