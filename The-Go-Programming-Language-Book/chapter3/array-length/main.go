package main

import "fmt"

func main() {
	// program to reverse an array
	arr := []int{1, 2, 3, 4}
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Println(arr[i])
	}

	var a uint8 = 255
	fmt.Println(a)
	a += 1
	fmt.Println(a)

}
