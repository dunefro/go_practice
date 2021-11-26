package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	start := time.Now()
	fmt.Println("Before", runtime.NumGoroutine())
	wg.Add(10)
	for i := 1; i <= 10; i++ {
		go func() {
			mu.Lock()
			for j := 1; j <= 5; j++ {
				// mu.Lock()
				// defer mu.Unlock()
				// fmt.Println("This is ", i)
				fmt.Println("foo", j)
			}
			mu.Unlock()
			wg.Done()
		}()
		// for j := 1; j <= 5; j++ {
		// 	// fmt.Println("This is ", i)
		// 	fmt.Println("foo", j)
		// }
	}
	// foo()
	bar()
	fmt.Println("Middle", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("After", runtime.NumGoroutine())
	fmt.Println("time", time.Since(start))
}

// func foo() {
// 	for i := 1; i <= 5; i++ {
// 		fmt.Println("foo", i)
// 	}
// 	// wg.Done()
// }

func bar() {
	for i := 1; i <= 5; i++ {
		fmt.Println("bar", i)
	}
}
