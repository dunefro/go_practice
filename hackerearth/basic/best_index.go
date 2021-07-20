package main

import "fmt"

func main() {
	var size, max int
	fmt.Scanln(&size)
	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	for i := 0; i < size; i++ {
		t := i
		k := i
		for j := 2; ; j++ {
			if t+j >= size {
				k = t
				break
			} else {
				t += j
			}
		}
		if sum := calculateSum(arr, i, k); max < sum {
			max = sum
		}
	}
	fmt.Println(max)
}
func calculateSum(arr []int, start, end int) int {
	var sum int
	for i := start; i <= end; i++ {
		sum += arr[i]
	}
	return sum
}
