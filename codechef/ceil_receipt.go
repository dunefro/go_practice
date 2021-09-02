package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// arr := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	T, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	for i := int64(0); i < T; i++ {
		var count int
		scanner.Scan()
		p, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		strconv.FormatInt(p, 2)
		if p > 2048 {
			for {
				if p <= 2048 {
					break
				} else {
					p = p - 2048
					count += 1
				}
			}
		}
		arr := strings.Split(strconv.FormatInt(p, 2), "")
		for i := 0; i < len(arr); i++ {
			if arr[i] == "1" {
				count += 1
			}
		}
		fmt.Println(count)
		// P, := int(p)

	}
}
