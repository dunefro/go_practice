package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	T, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	for i := int64(0); i < T; i++ {
		scanner.Scan()
		number, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		// factorial(number)
		fmt.Println(factorial(number))
	}
}
func factorial(n int64) *big.Int {
	// fact = 1
	fact := big.NewInt(1)
	// bign := big.NewInt(n)
	for i := int64(1); i <= n; i++ {
		bigi := big.NewInt(i)
		fact = fact.Mul(fact, bigi)
	}
	return fact
}
