package array

// Sum function for non-fixed size of array
func Sum(array []int) int {
	sum := 0
	for _, n := range array {
		sum += n
	}
	return sum
}

// Function to write the sum of all slices
func SumAll(slices ...[]int) (total int) {
	for _, slice := range slices {
		for _, n := range slice {
			total += n
		}
	}
	return
}
