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
		var N, X int64
		scanner.Scan()
		a1 := strings.Split(strings.TrimRight(scanner.Text(), " "), " ")
		for k := 0; k < len(a1); k++ {
			if k == 0 {
				N, _ = strconv.ParseInt(a1[k], 10, 64)
			} else {
				X, _ = strconv.ParseInt(a1[k], 10, 64)
			}
		}
		N = N + 1 - 1
		scanner.Scan()
		a2 := strings.Split(strings.TrimRight(scanner.Text(), " "), " ")
		var checker = make(map[int64]map[string]int)
		for j := 0; j < len(a2); j++ {
			P, _ := strconv.ParseInt(a2[j], 10, 64)
			var xorP int64
			xorP = int64(P ^ X)
			if len(checker[P]) == 0 {
				checker[P] = map[string]int{}
				checker[P]["countermain"] = 0
				checker[P]["counterxor"] = 0
			}
			checker[P]["countermain"] += 1
			if xorP != P {
				if len(checker[xorP]) == 0 {
					checker[xorP] = map[string]int{}
					checker[xorP]["countermain"] = 0
					checker[xorP]["counterxor"] = 0
				}
				checker[xorP]["counterxor"] += 1
			}
		}
		var maxval = -1
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
