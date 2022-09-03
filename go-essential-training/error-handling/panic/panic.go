package main

import "fmt"

func main() {
	vals := []int{10, 20, 30}
	/*
		n := vals[3] // it will cause panic
		fmt.Println(n)
	*/

	v, err := safeValue(vals, 3)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("v: %v", v)
}

func safeValue(vals []int, index int) (n int, err error) {

	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()

	return vals[index], nil

}
