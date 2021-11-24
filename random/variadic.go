package main

import "fmt"

func main() {
	// arr := [5]int{1, 2, 3, 4, 5}
	// foo(arr)
	foo(1, 2, 3, 4, 5)
}
func foo(a ...int) {
	sum := 0
	for i := range a {
		sum += i
	}
	fmt.Println(sum)
}
