package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	ch <- 42
	// Adding one more to the buffer without taking out to prev one
	fmt.Println(<-ch)
	ch <- 43
	fmt.Println(<-ch)
}
