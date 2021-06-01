package main

import "fmt"

// const is introduced so as to reduce creating the "hello " string again and again
const englishHello = "Hello "

func hello(name string) string {
	// Condition for empty string
	if name == "" {
		name = "World"
	}
	return englishHello + name + " !!"
}

func main() {
	fmt.Println(hello("Vedant"))
}
