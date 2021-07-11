package main

import (
	"fmt"
	"strconv"
)

func main() {
	var size int
	fmt.Scanln(&size)
	var arr = make([]int, size)
	// var lastDigit int
	var lastDigit string
	// str := ""
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &arr[i])
		// lastDigit = arr[i] % 10
		// str = str + strconv.Itoa(lastDigit)
	}
	// lastDigitno, _ := strconv.Atoi(str)
	// if lastDigitno%10 == 0 {
	// fmt.Println("Yes")
	// } else {
	// fmt.Println("No")
	// }
	// fmt.Printf("%T", arr[size-1])
	lastDigit = strconv.Itoa(arr[size-1])
	// fmt.Println(lastDigit
	if string(lastDigit[len(lastDigit)-1]) == "0" {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
