package array

import "testing"

func TestSum(t *testing.T) {

	t.Run("Running the test with fixed no of arrays", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if got != want {
			t.Errorf("Expected: %d Recieved: %d", got, want)
		}
	})
	t.Run("Running the test without fixed no of arrays", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		if got != want {
			t.Errorf("Expected: %d Output: %d", got, want)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Test to check the SumAll function", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{3, 5, 6, 2, 6, 2}, []int{3})
		want := 33
		if got != want {
			t.Errorf("Expected: %d Output: %d", got, want)
		}

	})
}
