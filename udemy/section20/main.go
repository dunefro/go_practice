package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	// Here we add one to the stack of routines
	wg.Add(1)
	// Here the routine leaves the main thread
	go foo()
	bar()
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	// WaitGroup will wait until all the routines are completed or
	// WaitGroup will wait until all the routines are removed from the stack
	wg.Wait()
	// To check no of routines after the wait group is executed fully
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
}
func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo", i)
	}
	// Here we define that we have completed the routine
	// This function then removes the routine from the stack
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar", 1)
	}
}
