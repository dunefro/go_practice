package main

import (
	"fmt"
	"math"
)

func main() {
	var size, k int
	fmt.Scanf("%d", &size)
	fmt.Scanf("%d", &k)
	var arr = make([]int, size)
	var tarr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	for i := 0; i < size; i++ {
		if arr[i] > 0 {
			rem := arr[i]
			fmt.Println("__________________", arr[i])
			for j := i + 1; j <= i+k; j++ {
				fmt.Println(arr[j])
				fmt.Println(rem)
				if arr[j] < 0 && rem > 0 {
					fmt.Println("Entered sinde")
					if rem >= int(math.Abs(float64(arr[j]))) {
						tarr[i] = tarr[i] - int(math.Abs(float64(arr[j])))
						tarr[j] = tarr[j] + int(math.Abs(float64(arr[j])))
						rem = rem - int(math.Abs(float64(arr[j])))
					}
				}
			}
		}
	}
	cost := 0
	for i := 0; i < size; i++ {
		cost = cost + int(math.Abs(float64(arr[i]+tarr[i])))
	}
	fmt.Println(cost)
}
