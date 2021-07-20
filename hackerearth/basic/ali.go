package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	var flag bool
	var i1, i2 int
	fmt.Scanf("%s", &str)
	for i := 0; i < len(str); i++ {
		if i == 2 {
			char := string(str[i])
			if char == "A" || char == "E" || char == "I" || char == "O" || char == "U" || char == "Y" {
				flag = true
				break
			}
		} else if i == 1 || i == 5 || i == 6 || i == 8 {
			continue
		} else {
			i1, _ = strconv.Atoi(string(str[i]))
			i2, _ = strconv.Atoi(string(str[i+1]))
			if (i1+i2)%2 != 0 {
				flag = true
				break
			}
		}
	}
	if flag {
		fmt.Println("invalid")
	} else {
		fmt.Println("valid")
	}
}
