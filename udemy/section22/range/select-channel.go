package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	q := make(chan int)
	go send(eve, odd, q)
	receive(eve, odd, q)
}
func send(e, o, q chan<- int) {
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}
func receive(e, o, q <-chan int) {
	// infinite loop
	for {
		// select will select one of three channels at random
		// q will be selected last because it is being added at the end of for loop
		select {
		case v := <-e:
			fmt.Println("This is eve channel", v)
		case v := <-o:
			fmt.Println("This is odd channel", v)
		case v := <-q:
			fmt.Println("This is quit channel", v)
			return // Closes the for loop
		}
	}
}
