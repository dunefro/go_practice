package array

// Sum function for non-fixed size of array
func Sum(array []int) int {
	sum := 0
	for _, n := range array {
		sum += n
	}
	return sum
}
