package main

import "fmt"

type Square struct {
	x      int
	y      int
	length int
}

func NewSquare(x int, y int, length int) (*Square, error) {
	if length <= 0 {
		return nil, fmt.Errorf("length must be greater than 0")
	}

	s := Square{
		x:      x,
		y:      y,
		length: length,
	}
	return &s, nil
}

func (s *Square) Move(dx int, dy int) {
	s.x += dx
	s.y += dy
}

func (s *Square) Area() int {
	return s.length * s.length
}

func main() {
	s, err := NewSquare(1, 1, 10)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		s.Move(2, 3)
		fmt.Printf("Square: %#v\n", s)
		fmt.Printf("Area: %v\n", s.Area())
	}
}
