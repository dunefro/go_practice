package main

import "fmt"

func main() {
	var str string
	fmt.Scanln(&str)
	countz := 0
	counto := 0
	// fmt.Println(str)
	for _, i := range str {
		if string(i) == "z" {
			countz += 1
		} else if string(i) == "o" {
			counto += 1
		} else {
			break
		}
	}
	if counto == 2*countz {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
