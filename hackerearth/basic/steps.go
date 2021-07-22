package main

import "fmt"

func main() {
	var size, steps int
	fmt.Scanf("%d", &size)
	var checkToContinue int
	var checkAll int
	var a = make([]int, size)
	var b = make([]int, size)
	min := 5000
	new_true_min := 5000
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &a[i])
		if min > a[i] {
			min = a[i]
		}
	}
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &b[i])
	}
	for {
		checkToContinue = 0
		checkAll = 0
		for i := 0; i < size; i++ {
			if a[i] != min && a[i] >= b[i] {
				steps += 1
				a[i] -= b[i]

				if min > a[i] {
					if a[i] < b[i] {
						if new_true_min == 5000 {
							new_true_min = a[i]
						}
						if new_true_min != a[i] {
							fmt.Println("-1")
							return
						}
					}
					min = a[i]
					checkAll = 0
				}
			} else if a[i] <= b[i] {
				checkToContinue += 1
			}
			if a[i] == min {
				checkAll += 1
			}
		}
		if checkAll == size {
			fmt.Println(steps)
			break
		} else if checkToContinue == size {
			fmt.Println("-1")
			break
		}
	}

}
