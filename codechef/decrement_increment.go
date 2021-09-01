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
	N, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	if N%4 == 0 {
		fmt.Println(N + int64(1))
	} else {
		fmt.Println(N - int64(1))
	}
}
