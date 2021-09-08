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
		var total int64
		scanner.Scan()
		arr := strings.Split(scanner.Text(), " ")
		N, _ := strconv.ParseInt(arr[0], 10, 64)
		A, _ := strconv.ParseInt(arr[1], 10, 64)
		B, _ := strconv.ParseInt(arr[2], 10, 64)
		scanner.Scan()
		passes := strings.Split(scanner.Text(), "")
		for i := int64(0); i < N; i++ {
			if passes[i] == "0" {
				total += A
			} else if passes[i] == "1" {
				total += B
			}
		}
		fmt.Println(total)
	}
}
