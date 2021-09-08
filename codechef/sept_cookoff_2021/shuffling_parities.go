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
	for k := int64(0); k < T; k++ {
		count := 0
		scanner.Scan()
		n, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		N := int(n)
		N = N + 1
		scanner.Scan()
		arr := strings.Split(scanner.Text(), " ")
		evenindexcount := 0
		oddindexcount := 0
		for i := 0; i < len(arr); i++ {
			Z, _ := strconv.ParseInt(arr[i], 10, 64)
			z := int(Z)
			if (z+i+1)%2 == 0 {
				if (i+1)%2 == 0 {
					evenindexcount += 1
				} else {
					oddindexcount += 1
				}
			} else {
				count += 1
			}
		}
		if evenindexcount > oddindexcount {
			count += oddindexcount * 2
		} else {
			count += evenindexcount * 2
		}
		fmt.Println(count)

	}
}
