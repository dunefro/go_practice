package main

import (
	"fmt"
	"strings"
)

func main() {
	var size int
	fmt.Scanln(&size)
	var house string
	fmt.Scanln(&house)
	if strings.Contains(house, "HH") {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
		fmt.Println(strings.ReplaceAll(house, ".", "B"))
	}
}
