package perimeter

import "testing"

func TestPermiter(t *testing.T) {

	checkgotandwant := func(t testing.TB, got float64, want float64) {
		if got != want {
			t.Errorf("Expected: %.2f Received: %.2f", got, want)
		}
	}

	t.Run("Testing the perimeter of rectangle", func(t *testing.T) {
		got := Permiter(10.0, 40.0)
		want := 100.0
		checkgotandwant(t, got, want)
	})
	t.Run("Testing the area of rectangle", func(t *testing.T) {
		got := Area(10.0, 40.0)
		want := 400.0
		checkgotandwant(t, got, want)
	})
}
