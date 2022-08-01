package main

import "fmt"

func main() {
	a := 10
	b := 20
	fmt.Printf("%d %[1]d %[2]d %[1]d", a, b)
}
