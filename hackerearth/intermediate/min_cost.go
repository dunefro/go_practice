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
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	cost := 0
	for i := 0; i < size; i++ {
		if arr[i] > 0 {
			rem := arr[i]
			for j := i + 1; j <= i+k; j++ {
				if arr[j] < 0 && rem > 0 {
					absValue := int(math.Abs(float64(arr[j])))
					if rem < absValue {
						absValue = rem
					}
					arr[i] = arr[i] - absValue
					arr[j] = arr[j] + absValue
					rem = rem - absValue
				}
			}
		}
		cost = cost + int(math.Abs(float64(arr[i])))
	}
	fmt.Println(cost)
}
