package main

import "fmt"

func main() {
	s := "Vedant 23"
	var name string
	var age int
	fmt.Sscanf(s, "%s %d", &name, &age)
	fmt.Printf("Name is %s and age is %d\n", name, age)
}
