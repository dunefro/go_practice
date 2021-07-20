package main

import "fmt"

func main() {
	var size, max, k, sum, arrSum int
	max = -1 << 63
	fmt.Scanln(&size)
	var arr = make([]int, size)
	var sumArr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &arr[i])
		arrSum += arr[i]
		sumArr[i] = arrSum
	}
	for i := 0; i < size; i++ {
		t := i
		for j := 2; ; j++ {
			if t+j >= size {
				k = t
				break
			} else {
				t += j
			}
		}
		if i == 0 {
			sum = sumArr[i]
		} else {
			sum = sumArr[k] - sumArr[i-1]
		}
		if max < sum {
			max = sum
		}
	}
	fmt.Println(max)
}
