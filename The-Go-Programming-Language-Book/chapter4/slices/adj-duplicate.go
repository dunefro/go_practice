package main

import "fmt"

func main() {
	s := []string{"a", "b", "cc", "cc", "d", "e", "e", "f", "f"}
	fmt.Println(removeAdjacent(s))
}

func removeAdjacent(s []string) []string {
	count := 0
	for i, v := range s {
		if i != len(s)-1 {
			if v == s[i+1] {
				if i+2 < len(s) {
					s[i+1] = s[i+2]
				}
				count++
			}
		}
	}
	return s[:count]
}
