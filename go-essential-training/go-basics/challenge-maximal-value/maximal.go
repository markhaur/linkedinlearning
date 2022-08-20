package main

import "fmt"

func main() {
	nums := []int{16, 8, 42, 4, 23, 15}

	max := nums[0]

	for _, value := range nums[1:] {
		if max < value {
			max = value
		}
	}

	fmt.Println(max)
}
