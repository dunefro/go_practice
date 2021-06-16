package main

import "fmt"

func main() {
	ch := make(chan int)
	// Send
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		// Closing the channel
		close(ch)

	}()
	// Receive
	for j := range ch {
		fmt.Println(j)
	}
}
