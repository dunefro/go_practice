package main

import "fmt"

func main() {
	ch := make(chan int)
	cs := make(chan<- int)
	cr := make(<-chan int)

	// Assigning general channel type to specific channel type
	// Works fine
	cs = ch
	cr = ch

	fmt.Printf("Send channel %T\n", cs)
	fmt.Printf("Receive channel %T\n", cr)

	// Assigning specific channel type to general channel type
	// Error
	ch = cr
	ch = cs
}
