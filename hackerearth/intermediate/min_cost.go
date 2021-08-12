package main

import "fmt"

func main() {
	var size, k int
	fmt.Scanf("%d", &size)
	fmt.Scanf("%d", &k)
	var arr = make([]int, size)
	var tarr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	for i:=0; i< size; i++ {
		if arr[i] > 0 {
			rem := arr[i]
			for j:=i+1; j<=i+k ; j++ {
				if arr[j] < 0 && rem > 0 {
					if tarr[i]
					tarr[i] = tarr[i] - math.Abs(arr[j]
					tarr[j] = tarr[j] + math.Abs(arr[j])
					rem = rem + arr[j]
				}
			}
		}
	}

}
