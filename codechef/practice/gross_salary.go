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
		s, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		salary := float64(s)
		var gross float64
		if salary < 1500.0 {
			gross = salary + (0.1 * salary) + (0.9 * salary)
		} else {
			gross = salary + 500.0 + (0.98 * salary)
		}
		fmt.Printf("%.2f\n", gross)
	}
}
