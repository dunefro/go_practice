package main

import (
	"fmt"
	"strings"
)

func ConvertToRoman(arabic int) string {

	var result strings.Builder

	romanmap := map[int]string{1: "I", 4: "IV", 5: "V", 9: "IX", 10: "X", 40: "XL", 50: "L", 90: "XC", 100: "C", 400: "CD", 500: "D", 900: "CM", 1000: "M"}
	no := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}

	for i := 0; i < len(no); i++ {
		if arabic == no[i] {
			result.WriteString(romanmap[no[i]])
			break
		} else if no[i] > arabic {
			fmt.Println(no[i], arabic)
			if no[i]-arabic > no[i-1] {
				result.WriteString(romanmap[no[i-1]] + ConvertToRoman(arabic-no[i-1]))
			} else {
				result.WriteString(ConvertToRoman(no[i]-arabic) + romanmap[no[i]])
			}
			break
		}
	}
	return result.String()
}
