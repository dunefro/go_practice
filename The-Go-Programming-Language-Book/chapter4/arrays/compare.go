package main

import "fmt"

func main() {
	a1 := [2]int{1, 2}
	a2 := [...]int{3, 4}
	a3 := [2]int{3, 4}
	fmt.Println(a1 == a2, a2 == a3, a1 == a3) // false true false
	// a4 := [3]int{3, 4}
	// fmt.Println(a3 == a4) // will give error
}
