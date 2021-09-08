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
		N, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		fmt.Println(factorial(N))
	}
}
func factorial(n int64) *big.Int {
	fact := big.NewInt(1)
	for i := int64(1); i <= n; i++ {
		I := big.NewInt(i)
		fact.Mul(fact, I)
	}
	return fact
}
