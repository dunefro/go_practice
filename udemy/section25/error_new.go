package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	// Expecting error
	_, err := sqrt(-10)
	if err != nil {
		fmt.Println(err.Error())
	}
	val, err := sqrt(64)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(val)
}
func sqrt(n int) (int, error) {
	// Returning error
	if n < 0 {
		return n, errors.New("Number should not be less then 0")
	}
	return int(math.Sqrt(float64(n))), nil
}
