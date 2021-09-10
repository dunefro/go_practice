package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	T, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	for i := int64(0); i < T; i++ {
		var N, X int
		scanner.Scan()
		a1 := strings.Split(strings.TrimRight(scanner.Text(), " "), " ")
		for k := 0; k < len(a1); k++ {
			if k == 0 {
				n, _ := strconv.ParseInt(a1[k], 10, 64)
				N = int(n)
			} else {
				x, _ := strconv.ParseInt(a1[k], 10, 64)
				X = int(x)
			}
		}
		N = N + 1 - 1
		scanner.Scan()
		a2 := strings.Split(scanner.Text(), " ")
		// fmt.Println(a2, N, X)
		var checker = make(map[int]map[string]int)
		for j := 0; j < len(a2); j++ {
			p, _ := strconv.ParseInt(a2[j], 10, 64)
			P := int(p)
			up := uint(P)
			ux := uint(X)
			xorP := int(up ^ ux)
			if len(checker[P]) == 0 {
				checker[P] = map[string]int{}
				checker[P]["countermain"] = 1
				checker[P]["counterxor"] = 0
			} else {
				checker[P]["countermain"] += 1
			}
			if xorP != P {
				if len(checker[xorP]) == 0 {
					checker[xorP] = map[string]int{}
					checker[xorP]["countermain"] = 0
					checker[xorP]["counterxor"] = 1
				} else {
					checker[xorP]["counterxor"] += 1
				}
			}
		}
		var maxval int
		var minop = int(math.MaxInt64)
		for _, v := range checker {
			if v["countermain"]+v["counterxor"] > maxval {
				maxval = v["countermain"] + v["counterxor"]
				minop = v["counterxor"]
			} else if v["countermain"]+v["counterxor"] == maxval {
				if v["counterxor"] < minop {
					minop = v["counterxor"]
				}
			}
		}
		fmt.Println(maxval, minop)
	}
}
