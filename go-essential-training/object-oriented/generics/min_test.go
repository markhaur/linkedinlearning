package generics_test

import (
	"fmt"

	"github.com/markhaur/linkedinlearning/go-essential-training/object-oriented/generics"
)

func ExampleMinArgs() {
	fmt.Println(generics.MinArgs(1, 2, 3, 4, 5, 6))
	fmt.Println(generics.MinArgs(1.0, 2.0, 3.0, 4.0, 5.0, 6.0))
	fmt.Println(generics.MinArgs("joe", "billy", "perry", "biden", "lilly"))
	fmt.Println(generics.MinArgs[int]())
	fmt.Println(generics.MinArgs[string]())
	// Output: 1 <nil>
	// 1 <nil>
	// biden <nil>
	// 0 empty slice
	//  empty slice
}

func ExampleMinArray() {
	fmt.Println(generics.MinArray([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(generics.MinArray([]float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0}))
	fmt.Println(generics.MinArray([]string{"joe", "billy", "perry", "biden", "lilly"}))
	fmt.Println(generics.MinArray([]int{}))
	// Output: 1 <nil>
	// 1 <nil>
	// biden <nil>
	// 0 empty slice
}
