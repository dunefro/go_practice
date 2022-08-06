package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Print("Enter your number: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	fmt.Println(s)
	n, _ := strconv.ParseInt(s, 10, 64)
	fmt.Println(n * 2)
}
