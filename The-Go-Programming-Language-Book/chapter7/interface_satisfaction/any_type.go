package main

import "fmt"

func main() {
	var any interface{}

	any = 2 // int
	fmt.Println(any)
	any = true // bool
	fmt.Println(any)
	any = []int{1, 2, 3} // []int
	fmt.Println(any)

}
