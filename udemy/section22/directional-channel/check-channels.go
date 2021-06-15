package main

import "fmt"

func main() {
	ch := make(chan int)
	cs := make(chan<- int)
	cr := make(<-chan int)

	ch <- 42          // Allowed
	fmt.Println(<-ch) // Allowed

	cs <- 43          // Allowed
	fmt.Println(<-cs) // Not Allowed

	cr <- 44          // Not allowed
	fmt.Println(<-cr) // Allowed
}
