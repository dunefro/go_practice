package main

import "fmt"

func main() {
	ch := make(chan int)
	cs := make(chan<- int)
	cr := make(<-chan int)

	// Allowed
	fmt.Printf("Channel %T\n", (chan<- int)(ch))
	fmt.Printf("channel %T\n", (<-chan int)(ch))

	// Not Allowed
	fmt.Printf("Send channel %T\n", (chan int)(cs))
	fmt.Printf("Receieve channel %T\n", (chan int)(cr))
}
