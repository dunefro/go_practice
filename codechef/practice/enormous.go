package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, count int
	var x, k int64
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &k)
	for i := 0; i < n; i++ {
		scanner.Scan()
		x, _ = strconv.ParseInt(scanner.Text(), 10, 64)

		if x%k == 0 {
			count += 1
		}
	}
	fmt.Println(count)
}
