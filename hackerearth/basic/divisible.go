package main

import (
	"fmt"
	"strconv"
)

func main() {
	var size int
	fmt.Scanln(&size)
	var no string
	// var final strings.Builder
	var result int
	var temp int
	for i := 0; i < size; i++ {
		fmt.Scanf("%s", &no)
		if i < size/2 {
			temp, _ = strconv.Atoi(string(no[0]))
		} else {
			temp, _ = strconv.Atoi(string(no[len(no)-1]))
		}
		fmt.Println("for ", i, "value is", temp)
		if i%2 == 0 {
			result += temp
		} else {
			result -= temp
		}
	}
	if result%11 == 0 {
		fmt.Println("OUI")
	} else {
		fmt.Println("NON")
	}
}
