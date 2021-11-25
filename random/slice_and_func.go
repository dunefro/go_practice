package main

import "fmt"

func main() {
	arr := [4]int{1, 2, 3, 4}
	fmt.Println(arr)
	foo(&arr)
	fmt.Println(arr)
}
func foo(arr *[4]int) {
	(*arr)[0] = -1
	fmt.Println(*arr)
}
