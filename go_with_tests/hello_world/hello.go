package main

import "fmt"

// const is introduced so as to reduce creating the "hello " string again and again
const englishHello = "Hello "
const hindiHello = "Namaste "
const spanishHello = "Hola "

func Hello(name string, language string) string {
	// Condition for empty string
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name + " !!"
}

// (prefix string) -> creates a string variable prefix and automatically return it at the end by simple return statement
func greetingPrefix(language string) (prefix string) {

	switch language {
	case "hindi":
		prefix = hindiHello
	case "spanish":
		prefix = spanishHello
	default:
		prefix = englishHello
	}
	return
}

func main() {
	fmt.Println(Hello("", "hindi"))
	fmt.Println(Hello("Vedant", ""))

}
