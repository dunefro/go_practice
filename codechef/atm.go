package main

import (
	"fmt"
	"math"
)

func main() {
	var withdrawl float64
	var total float64
	var rem float64
	fmt.Scanf("%f", &withdrawl)
	fmt.Scanf("%f", &total)
	if withdrawl <= total-0.50 {
		if math.Mod(withdrawl, 5) == 0 {
			rem = total - (withdrawl + 0.50)
		} else {
			rem = total
		}
	} else {
		rem = total
	}
	fmt.Printf("%.2f", rem)
}
