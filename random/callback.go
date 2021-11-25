package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(evensum(sum, x...))
}

func sum(x ...int) int {
	total := 0
	fmt.Println(x)
	for _, v := range x {
		total += v
	}
	fmt.Println(total)
	return total
}

func evensum(f func(...int) int, x ...int) int {
	var evearry []int
	for _, v := range x {
		if v%2 == 0 {
			evearry = append(evearry, v)
		}
	}
	fmt.Println(evearry)
	return f(evearry...)
}
