package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	T, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	arr := make([]int, T)
	for i := int64(0); i < T; i++ {
		scanner.Scan()
		n, _ := (strconv.ParseInt(scanner.Text(), 10, 64))
		arr[i] = int(n)
	}
	sort.Ints(arr)
	for i := int64(0); i < T; i++ {
		fmt.Println(arr[i])
	}
}
