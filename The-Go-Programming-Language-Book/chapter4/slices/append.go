package main

import "fmt"

func main() {
	var x = make([]int, 3, 3)
	x[0] = 1
	x[1] = 2
	x[2] = 3
	fmt.Println(appendInt(x[:3], 4))

}
func appendInt(x []int, y int) []int {
	if cap(x) >= len(x)+1 {
		fmt.Println("orginal")
		x = append(x, y)
		return x
	} else {
		z := make([]int, len(x)+1, len(x)+1)
		copy(z, x)
		z[len(x)] = y
		return z
	}
}
