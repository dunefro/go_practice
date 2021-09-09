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
		D, _ := strconv.ParseInt(arr[0], 10, 64)
		N, _ := strconv.ParseInt(arr[1], 10, 64)
		sum := int64(0)
		for j := int64(0); j < D; j++ {
			sum = (N * (N + 1)) / 2
			N = sum
		}
		fmt.Println(sum)
	}
}
