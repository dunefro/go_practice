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
		hardness, _ := strconv.ParseInt(arr[0], 10, 64)
		tensile, _ := strconv.ParseInt(arr[2], 10, 64)
		carbon, _ := strconv.ParseFloat(arr[1], 32)
		var grade int
		if hardness > 50 {
			if carbon < 0.7 {
				if tensile > 5600 {
					grade = 10
				} else {
					grade = 9
				}
			} else if tensile > 5600 {
				grade = 7
			} else {
				grade = 6
			}
		} else {
			if carbon < 0.7 {
				if tensile > 5600 {
					grade = 8
				} else {
					grade = 6
				}
			} else {
				if tensile > 5600 {
					grade = 6
				} else {
					grade = 5
				}
			}
		}
		fmt.Println(grade)
	}
}
