package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Print("Please enter an index: ")
	var index int
	fmt.Scanf("%d", &index)
	element := arr[index]
	fmt.Println("Element at index", index, "is:", element)
}
