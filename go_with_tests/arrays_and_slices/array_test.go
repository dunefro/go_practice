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
