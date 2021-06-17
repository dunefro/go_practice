package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		ch <- 24
		close(ch)
	}()
	v, ok := <-ch
	fmt.Println(v, ok)

	// Printing it second time
	v, ok = <-ch
	fmt.Println(v, ok)
}
