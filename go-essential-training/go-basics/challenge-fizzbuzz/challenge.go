package main

import "fmt"

func main() {

	for i := 1; i <= 20; i++ {

		if i%3 == 0 {
			fmt.Print("fizz ")
		}

		if i%5 == 0 {
			fmt.Print("buzz")
		}

		if i%3 != 0 && i%5 != 0 {
			fmt.Print(i)
		}

		fmt.Println("")
	}
}
