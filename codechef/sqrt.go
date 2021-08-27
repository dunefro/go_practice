package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	T, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	for i := int64(0); i < T; i++ {
		scanner.Scan()
		n, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		fmt.Println(int(math.Sqrt(float64(n)) + 0.5))
	}
}
