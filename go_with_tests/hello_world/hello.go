package main

import "fmt"

// const is introduced so as to reduce creating the "hello " string again and again
const englishHello = "Hello "
const hindiHello = "Namaste "

func hello(name string, language string) string {
	// Condition for empty string
	if name == "" {
		name = "World"
	}
	if language == "hindi" {
		return hindiHello + name + " !!"
	}

	return englishHello + name + " !!"
}

func main() {
	fmt.Println(hello("Vedant", "hindi"))
}
