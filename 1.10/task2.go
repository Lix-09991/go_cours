package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func PrintAreas(shapes []Shape) {
	for i, s := range shapes {
		fmt.Printf("Фигура %d: площадь = %.2f\n", i+1, s.Area())
	}
}

func main() {
	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 4, Height: 6},
		Circle{Radius: 2.5},
	}

	PrintAreas(shapes)
}
