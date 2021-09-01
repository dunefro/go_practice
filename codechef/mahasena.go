package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var lucky_count, unlucky_count int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	N := int(n)
	soldiers := make([]int, n)
	scanner.Scan()
	arr := strings.Split(scanner.Text(), " ")
	for i := 0; i < N; i++ {
		temp, _ := strconv.ParseInt(arr[i], 10, 64)
		soldiers[i] = int(temp)
		if soldiers[i]%2 == 0 {
			lucky_count += 1
		} else {
			unlucky_count += 1
		}
	}
	if lucky_count > unlucky_count {
		fmt.Println("READY FOR BATTLE")
	} else {
		fmt.Println("NOT READY")
	}
}
