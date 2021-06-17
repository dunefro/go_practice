package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(eve, odd, quit)
	receive(eve, odd, quit)
}
func send(e, o, q chan<- int) {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(q)
}
func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("Number is eve", v)
		case v := <-o:
			fmt.Println("Number is odd", v)
		case v := <-q:
			fmt.Println("Numbers quits", v)
			return

		}
	}
}
