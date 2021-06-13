package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	const n = 10
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(n)
	counter := 0
	for i := 1; i <= n; i++ {
		go func() {
			// We are increasing the counter independently from i
			// So we declare another variable v
			mu.Lock()
			v := counter
			v++
			runtime.Gosched()
			counter = v
			wg.Done()
			mu.Unlock()
		}()
	}
	fmt.Println("NO of go routines:", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println(counter)
}
