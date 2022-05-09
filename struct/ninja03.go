package _struct

import "fmt"

func Ninja03() {

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

	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "white",
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "white",
		},
		luxury: true,
	}

	fmt.Printf("%v", t)
	fmt.Printf("%v", s)
}
