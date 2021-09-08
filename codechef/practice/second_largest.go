package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
		narr := make([]int, 3)
		for j := 0; j < 3; j++ {
			number, _ := strconv.Atoi(arr[j])
			narr[j] = number
		}
		sort.Ints(narr)
		fmt.Println(narr[1])
	}
}
