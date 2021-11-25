package main

import "fmt"

func main() {
	func() {
		fmt.Println("This is annonymous function")
	}()
	func(x int) {
		fmt.Println("The value of the variable is", x)
	}(24)
}
