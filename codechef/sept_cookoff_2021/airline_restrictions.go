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
		A, _ := strconv.ParseInt(arr[0], 10, 64)
		B, _ := strconv.ParseInt(arr[1], 10, 64)
		C, _ := strconv.ParseInt(arr[2], 10, 64)
		D, _ := strconv.ParseInt(arr[3], 10, 64)
		E, _ := strconv.ParseInt(arr[4], 10, 64)
		sumAB := A + B
		sumBC := B + C
		sumAC := A + C
		var flag string
		if (sumAB <= D && C <= E) || (sumAB <= E && C <= D) {
			flag = "YES"
		} else if (sumBC <= D && A <= E) || (sumBC <= E && A <= D) {
			flag = "YES"
		} else if (sumAC <= D && B <= E) || (sumAC <= E && B <= D) {
			flag = "YES"
		} else {
			flag = "NO"
		}
		fmt.Println(flag)

	}
}
