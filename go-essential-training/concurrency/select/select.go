package main

import "fmt"

func main() {
	ch1, ch2 := make(chan int), make(chan int)

	go func() {
		ch1 <- 32
	}()

	select {
	case val := <-ch1:
		fmt.Printf("got %d from ch1\n", val)
	case val := <-ch2:
		fmt.Printf("got %d from ch2\n", val)
	}

	// more work to do
}
