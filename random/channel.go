package main

import "fmt"

func main() {
	fmt.Println("Main Begins")
	ch := make(chan int)
	// Caliing go routine here
	go printChannelData(ch)
	// Passing the value to channel
	ch <- 42
	fmt.Println("Main ended")
}

func printChannelData(c chan int) {
	fmt.Println("Accepting the channel with the value", <-c)
}
