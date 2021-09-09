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
		var total float64
		arr := strings.Split(scanner.Text(), " ")
		Q, _ := strconv.ParseFloat(arr[0], 32)
		P, _ := strconv.ParseFloat(arr[1], 32)
		mul := Q * P
		if Q > 1000 {
			total = mul - (mul * 0.1)
		} else {
			total = mul
		}
		fmt.Printf("%.6f\n", total)

	}
}
