package main

import "fmt"

func main() {
	// anonymous function declared
	sum := func(a, b int) int {
		return a + b
	}(4, 3) // parameters passed to the function
	fmt.Println(sum)
}
