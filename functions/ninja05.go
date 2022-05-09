package functions

import (
	"fmt"
	"math"
)

type ninja05Shape interface {
	area() float64
}

type ninja05Square struct {
	length float64
}

type ninja05Circle struct {
	radius float64
}

func (s ninja05Square) area() float64 {
	return s.length * s.length
}

func (s ninja05Circle) area() float64 {
	return s.radius * s.radius * math.Pi
}

func ninja05SayArea(s ninja05Shape) {
	fmt.Printf("The area is %v\n", s.area())
}

func Ninja05() {
	ninja05SayArea(
		ninja05Circle{
			radius: 12.345,
		},
	)
	ninja05SayArea(
		ninja05Square{
			length: 15,
		},
	)
}
