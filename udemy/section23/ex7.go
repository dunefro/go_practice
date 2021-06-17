package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(10)
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				ch <- j
			}
		}()
		wg.Done()
	}
	receive(ch)
	wg.Wait()
}
func receive(ch <-chan int) {
	for i := 0; i < 100; i++ {
		fmt.Println(<-ch)
	}
}
