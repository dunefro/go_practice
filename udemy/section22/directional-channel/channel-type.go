package main

import "fmt"

func main() {
	ch := make(chan int)
	cs := make(chan<- int)
	cr := make(<-chan int)
	fmt.Printf("Channel %T\n", ch)
	fmt.Printf("Send channel %T\n", cs)
	fmt.Printf("Receive channel %T\n", cr)
}
