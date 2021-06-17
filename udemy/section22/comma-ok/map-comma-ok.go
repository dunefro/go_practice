package main

import "fmt"

func main() {
	m := map[string]int{
		"key1": 1,
		"key2": 2,
	}
	// To print a value that exists
	v, ok := m["key1"]
	fmt.Println(v, ok)

	// To print a value that doesn't exists
	v, ok = m["somerandomkey"]
	fmt.Println(v, ok)
}
