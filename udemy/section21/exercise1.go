package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("This is the first go routine")
		wg.Done()
	}()
	go func() {
		fmt.Println("This is the second go routine")
		wg.Done()
	}()
	wg.Wait()
}
