package main

import "fmt"

func main() {
	book := "The colours of magic"
	fmt.Println(book)

	fmt.Println(len(book))

	fmt.Printf("book[0] = %v, Type = %T\n", book[0], book[0])
	// uint8 is a byte

	// strings in go are immutable
	//book[0] = 72

	// Slice (start, end), 0 based, half empty range
	fmt.Println(book[4:11])

	// Slice, no end
	fmt.Println(book[4:])

	// Slice, no start
	fmt.Println(book[:4])

	// Use + to concatenate strings
	fmt.Println("t" + book[1:])

	// Unicode
	fmt.Println("It was ½ price")

	// Multiline
	poem := `
		The road goes ever on
		Down from the door where it began
		...
	`
	fmt.Println(poem)
}
