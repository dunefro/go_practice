package main

import (
	"fmt"
	"strconv"
)

func main() {
	var size int
	fmt.Scanln(&size)
	var arr = make([]int, size
	var lastDigit string
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	lastDigit = strconv.Itoa(arr[size-1])
	if string(lastDigit[len(lastDigit)-1]) == "0" {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
