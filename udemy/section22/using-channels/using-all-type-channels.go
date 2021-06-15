package main

import "fmt"

func main() {
	c := make(chan int)

	// Send
	go foo(c)
	// Receive
	bar(c)
	fmt.Println("Exit")
}

// values are passed as a reference in channel
// Send
func foo(cs chan<- int) {
	cs <- 42
}

// Receive
func bar(cr <-chan int) {
	fmt.Println(<-cr)
}
