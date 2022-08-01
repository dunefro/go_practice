package main

import (
	"cyrpto/rand"
	"fmt"
)

func main() {
	// x := 5
	// if y := x*2; y == 12 {

	// }
	v, _ := rand.Intn(rand.Reader, 9)
	fmt.Println(v)
}
