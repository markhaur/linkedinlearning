package main

import (
	"fmt"
	"math"
)

func main() {

	square, err := sqrt(2.0)

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	} else {
		fmt.Printf("Square: %v\n", square)
	}

	square, err = sqrt(-2.0)

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	} else {
		fmt.Printf("Square: %v\n", square)
	}
}

func sqrt(num float64) (float64, error) {

	if num < 0 {
		return 0.0, fmt.Errorf("can't sqrt of negative value %v", num)
	}
	return math.Sqrt(num), nil
}
