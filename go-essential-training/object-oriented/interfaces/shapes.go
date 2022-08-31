package main

import (
	"fmt"
	"math"
)

type Square struct {
	Length float64
}

func (s Square) Area() float64 {
	return s.Length * s.Length
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func sumAreas(shapes []Shape) float64 {
	var totalArea float64

	for _, shape := range shapes {
		totalArea += shape.Area()
	}

	return totalArea
}

type Shape interface {
	Area() float64
}

func main() {

	s := Square{Length: 20}
	fmt.Println(s.Area())

	c := Circle{Radius: 10}
	fmt.Println(c.Area())

	shapes := []Shape{s, c}

	fmt.Println(sumAreas(shapes))
}
