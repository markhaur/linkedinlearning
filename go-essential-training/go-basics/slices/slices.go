package main

import "fmt"

func main() {

	loons := []string{"bugs", "daffy", "taz"}
	fmt.Printf("loons = %v, type = %T\n", loons, loons)

	// Length
	fmt.Printf("len(loons)=%v\n", len(loons))

	fmt.Println("-------------------")
	// 0 indexing
	fmt.Printf("loons[0]=%v\n", loons[1])

	fmt.Println("-------------------")
	// slices
	fmt.Printf("loons[1:] = %v\n", loons[1:])

	fmt.Println("-------------------")
	for i := 0; i < len(loons); i++ {
		fmt.Printf("loons[%v] = %v\n", i, loons[i])
	}

	fmt.Println("-------------------")
	// single value range

	for i := range loons {
		fmt.Printf("loons[%v] = %v\n", i, loons[i])
	}

	fmt.Println("-------------------")
	// double value range

	for i, name := range loons {
		fmt.Printf("index = %v, name = %v\n", i, name)
	}

	fmt.Println("-------------------")
	// double value range, ignoring index by using _

	for _, name := range loons {
		fmt.Printf("only name = %v\n", name)
	}

	fmt.Println("-------------------")
	// append
	loons = append(loons, "elmer")
	fmt.Printf("loons = %v\n", loons)

}
