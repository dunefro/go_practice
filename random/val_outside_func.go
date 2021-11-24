package main

import "fmt"

var y = 24

func main() {
	fmt.Println("Value of y in main is", y)
	foo()
	bar()

}
func foo() {
	y = y + 10
	fmt.Println("Value of y in foo is", y)
}

func bar() {
	y = y + 10
	fmt.Println("Value of y in bar is", y)
}
