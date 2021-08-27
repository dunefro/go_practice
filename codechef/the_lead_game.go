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
	W := int64(0)
	L := int64(0)
	sumP1 := int64(0)
	sumP2 := int64(0)
	for i := int64(0); i < T; i++ {
		scanner.Scan()
		arr := strings.Split(scanner.Text(), " ")
		P1, _ := strconv.ParseInt(arr[0], 10, 64)
		P2, _ := strconv.ParseInt(arr[1], 10, 64)
		sumP1 += P1
		sumP2 += P2
		if sumP1 > sumP2 {
			if L < sumP1-sumP2 {
				W = 1
				L = sumP1 - sumP2
			}
		} else if sumP2 > sumP1 {
			if L < sumP2-sumP1 {
				W = 2
				L = sumP2 - sumP1
			}
		}
	}
	fmt.Println(W, L)
}
