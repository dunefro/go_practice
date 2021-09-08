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
		n, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		N := int(n)
		str_N := strconv.Itoa(N)
		rev_str_N := reverseString(str_N)
		rev_N, _ := strconv.Atoi(rev_str_N)
		if N == rev_N {
			fmt.Println("wins")
		} else {
			fmt.Println("loses")
		}

	}
}
func reverseString(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-1-i] = runes[len(runes)-1-i], runes[i]
	}
	return string(runes)
}
