package main

import "fmt"

func main() {
	var testcase, cp1, cp2, cost1, cost2, n, p1, p2, total int
	fmt.Scanf("%d", &testcase)
	for i := 0; i < testcase; i++ {
		cp1 = 0
		cp2 = 0
		total = 0
		fmt.Scanf("%d %d", &cost1, &cost2)
		fmt.Scanf("%d", &n)
		for j := 0; j < n; j++ {
			fmt.Scanf("%d %d", &p1, &p2)
			if p1 == 1 {
				cp1 += 1
			}
			if p2 == 1 {
				cp2 += 1
			}
		}
		if cost1 > cost2 {
			if cp1 > cp2 {
				total = cp2*cost1 + cp1*cost2
			} else {
				total = cp1*cost1 + cp2*cost2
			}
		} else {
			if cp1 > cp2 {
				total = cp1*cost1 + cp2*cost2
			} else {
				total = cp2*cost1 + cp1*cost2
			}
		}
		fmt.Println(total)
	}
}
