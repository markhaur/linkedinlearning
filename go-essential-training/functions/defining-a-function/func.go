package main

import "fmt"

func main() {

	addResult := add(1, 1)
	fmt.Printf("addResult: %v\n", addResult)

	div, mod := divmod(7, 3)
	fmt.Printf("Div: %v, Mod: %v\n", div, mod)

}

func add(a int, b int) int {
	return a + b
}

func divmod(a int, b int) (int, int) {
	return a / b, a % b
}
