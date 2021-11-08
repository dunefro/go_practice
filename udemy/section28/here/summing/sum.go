// Package summing is used for sum operations
package summing

// Summing functions take variadic parameters to calculate sum of integers
func Summing(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
