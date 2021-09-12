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
	T, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	for i := int64(0); i < T; i++ {
		checker := make(map[string]map[string]int)
		scanner.Scan()
		text := scanner.Text()
		textarr := []rune(text)
		for j, k := 0, len(textarr)-1; j < len(textarr)/2; j, k = j+1, k-1 {
			if len(checker[string(textarr[j])]) == 0 {
				checker[string(textarr[j])] = map[string]int{}
				checker[string(textarr[j])]["leftcount"] = 0
				checker[string(textarr[j])]["rightcount"] = 0
			}
			checker[string(textarr[j])]["leftcount"] += 1
			if len(checker[string(textarr[k])]) == 0 {
				checker[string(textarr[k])] = map[string]int{}
				checker[string(textarr[k])]["leftcount"] = 0
				checker[string(textarr[k])]["rightcount"] = 0
			}
			checker[string(textarr[k])]["rightcount"] += 1
		}
		output := "YES"
		for _, v := range checker {
			if v["leftcount"] != v["rightcount"] {
				output = "NO"
				break
			}
		}
		fmt.Println(output)
	}
}
