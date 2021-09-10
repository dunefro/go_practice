package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	T, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	for i := int64(0); i < T; i++ {
		min := math.MaxInt64
		secondmin := math.MaxInt64
		scanner.Scan()
		N, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		var intarr = make([]int, N)
		scanner.Scan()
		arr := strings.Split(scanner.Text(), " ")
		for j := 0; j < len(arr); j++ {
			a, _ := strconv.ParseInt(arr[j], 10, 64)
			intarr[j] = int(a)
		}
		for j := 0; j < len(intarr); j++ {
			a := intarr[j]
			if a <= min {
				secondmin = min
				min = a
			} else if a < secondmin {
				secondmin = a
			}
		}
		fmt.Println(min + secondmin)
		// fmt.Println(min, secondmin)
	}
}
