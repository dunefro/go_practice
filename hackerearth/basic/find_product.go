package main

import (
	"fmt"
	"math"
)

func main() {
	var size int
	fmt.Scanln(&size)
	var arr = make([]int, size)
	product := 1
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &arr[i])
		product = (product * arr[i]) % (int(math.Pow(10, 9)) + 7)
	}
	fmt.Println(product)
}
