package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	T, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	for i := int64(0); i < T; i++ {
		scanner.Scan()
		B, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		B = B - 2
		B = B / 2
		fmt.Println(B * (B + 1) / 2)
	}
}
