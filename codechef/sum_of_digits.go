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
		arr := strings.Split(scanner.Text(), "")
		// number, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		var total int64
		for _, number := range arr {
			numberInInt, _ := strconv.ParseInt(number, 10, 64)
			total += numberInInt
		}
		fmt.Println(total)

	}

}
