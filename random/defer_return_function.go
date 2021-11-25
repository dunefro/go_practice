package main

import "fmt"

func main() {
	foo()()
}

func foo() func() {
	defer func() {
		fmt.Println("I am defered")
	}()
	return func() {
		fmt.Println("I am returned")
	}
}
