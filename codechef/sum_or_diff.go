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
	N1, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	scanner.Scan()
	N2, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	if N1 > N2 {
		fmt.Println(N1 - N2)
	} else {
		fmt.Println(N1 + N2)
	}
}
