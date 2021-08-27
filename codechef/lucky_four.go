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
		arr := strings.Split(scanner.Text(), "")
		count := 0
		for _, n := range arr {
			if n == "4" {
				count += 1
			}
		}
		fmt.Println(count)
	}
}
