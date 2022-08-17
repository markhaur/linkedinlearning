package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 3; i++ {
		fmt.Printf("i = %v\n", i)
	}

	fmt.Println("---------- Breaking a for loop in the middle ----------")

	for i := 0; i < 3; i++ {
		if i > 1 {
			break
		}
		fmt.Printf("printing i = %v if loop is not breaked\n", i)
	}

	fmt.Println("---------- Using continue in a for loop ----------")

	for i := 0; i < 3; i++ {
		if i < 1 {
			continue
		}
		fmt.Printf("printing i = %v if continue statement is not executed\n", i)
	}

	fmt.Println("---------- Using for loop as while in other languages ----------")

	a := 0
	for a < 3 {
		fmt.Printf("a = %v\n", a)
		a++
	}

	fmt.Println("---------- Using for loop as infinite loop ----------")

	b := 0
	for {
		if b > 1 {
			break
		}
		fmt.Printf("value of b = %v if infinite loop not break\n", b)
		b++
	}

}
