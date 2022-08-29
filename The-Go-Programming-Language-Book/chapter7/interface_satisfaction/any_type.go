package main

import "fmt"

func main() {
	var any interface{}

	any = 2 //int
	fmt.Println(any)
	any = true
	fmt.Println(any)
	any = []int{1, 2, 3}
	fmt.Println(any)

}
