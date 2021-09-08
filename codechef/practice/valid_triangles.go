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
		angle1, _ := strconv.ParseInt(arr[0], 10, 64)
		angle2, _ := strconv.ParseInt(arr[1], 10, 64)
		angle3, _ := strconv.ParseInt(arr[2], 10, 64)
		if angle1+angle2+angle3 == int64(180) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
