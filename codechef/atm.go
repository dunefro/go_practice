package main

import (
	"fmt"
	"math"
)

func main() {
	var w float64
	var total float64
	var rem float64
	fmt.Scanf("%f", &w)
	fmt.Scanf("%f", &total)
	if w <= total-0.50 {
		if math.Mod(w, 5) == 0 {
			rem = total - (w + 0.50)
		} else {
			rem = total
		}
	} else {
		rem = total
	}
	fmt.Printf("%.2f", rem)
}
