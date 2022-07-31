package main

import "fmt"

func main() {
	a := foo()
	fmt.Println(*a)
}

func foo() *int {
	v := 1
	return &v
}
