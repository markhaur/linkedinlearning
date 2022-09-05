package main

import "fmt"

func main() {
	ch := make(chan int)
	// ch <- 355 // sending on channel. it will block the main go routine

	go func() {
		ch <- 335
	}()

	val := <-ch
	fmt.Printf("got %d\n", val)

	fmt.Println("--------")
	counter := 3
	go func() {
		for i := 0; i < counter; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i
		}
	}()

	for i := 0; i < counter; i++ {
		val := <-ch
		fmt.Printf("receiving %d\n", val)
	}

	fmt.Printf("----------\n")

	go func() {
		for i := 0; i < counter; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i
		}
		close(ch)
	}()

	for i := range ch {
		fmt.Printf("receiving: %d\n", i)
	}

	// buffered channel
	bch := make(chan int, 1)
	bch <- 3
	fmt.Printf("bch rec: %d\n", <-bch)
	close(bch)
}
