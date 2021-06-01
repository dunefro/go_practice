package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutine:", runtime.NumGoroutine())

	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex
	for i := 0; i < gs; i++ {
		go func() {
			// Locking the variables here till it mutext unlocks
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
			// Unlocking the variables from lock till here
			mu.Unlock()
		}()
		fmt.Println("Goroutine:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println(counter)
}
