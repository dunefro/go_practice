package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	T, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	for i := int64(0); i < T; i++ {
		scanner.Scan()
		arr := strings.Split(scanner.Text(), " ")
		A, _ := strconv.ParseInt(arr[0], 10, 64)
		B, _ := strconv.ParseInt(arr[1], 10, 64)
		if A > B {
			fmt.Println(">")
		} else if A < B {
			fmt.Println("<")
		} else {
			fmt.Println("=")
		}

	}
}
