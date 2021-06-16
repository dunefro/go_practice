package main

import "fmt"

func main() {
	ch := make(chan int)
	// Send
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()
	// Receive
	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}
}
