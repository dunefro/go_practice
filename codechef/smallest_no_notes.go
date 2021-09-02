package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	arr := []int64{100, 50, 10, 5, 2, 1}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	T, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	for i := int64(0); i < T; i++ {
		count := int64(0)
		scanner.Scan()
		N, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		for j := 0; j < 6; j++ {
			quotient := N / arr[j]
			count += quotient
			rem := N % arr[j]
			if rem == 0 {
				break
			} else {
				N = rem
			}
		}
		fmt.Println(count)
	}
}
