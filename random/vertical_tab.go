package main

import (
	"fmt"
	"time"
)

func main() {
	var x = 10
	for i := 0; i < x; i++ {
		op := ""
		for j := 0; j < i*10; j++ {
			op += "#"
		}
		fmt.Printf("%s", op)
		fmt.Printf("\r")
		time.Sleep(1 * time.Second)
	}
	fmt.Println()

}
