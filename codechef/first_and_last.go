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
		number := scanner.Text()
		arr := strings.Split(number, "")
		first_digit, _ := strconv.ParseInt(arr[0], 10, 64)
		last_digit, _ := strconv.ParseInt(arr[len(arr)-1], 10, 64)
		fmt.Println(first_digit + last_digit)
	}
}
