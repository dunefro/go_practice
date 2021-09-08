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
		N, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		if N == int64(1) {
			fmt.Println("no")
			continue
		}
		var check bool
		for j := int64(2); j <= N/2; j++ {
			if N%j == 0 {
				check = true
				break
			}
		}
		if check {
			fmt.Println("no")
		} else {
			fmt.Println("yes")
		}
	}
}
