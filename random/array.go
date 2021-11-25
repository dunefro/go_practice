package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	s := arr[:]
	fmt.Println(s)
	s[2] = 5
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	s[3] = 6
	// s = append(s, 6)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}
