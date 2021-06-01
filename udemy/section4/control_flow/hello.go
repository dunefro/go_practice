package main

import "fmt"

func main() {
	fmt.Println("Hello World this is where my program starts")
	foo()
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
func foo() {
	fmt.Println("This is being called from the function foo. How do you feel my boy I think you mmust have been feeling really good, ins;t it")
}
