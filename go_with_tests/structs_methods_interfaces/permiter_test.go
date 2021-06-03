package perimeter

import "testing"

type Rectangle struct {
	width   float64
	breadth float64
}

type Circle struct {
	radius float64
}

func TestPermiter(t *testing.T) {

	checkgotandwant := func(t testing.TB, got float64, want float64) {
		if got != want {
			t.Errorf("Expected: %g Received: %g", got, want)
		}
	}

	t.Run("Testing the perimeter of rectangle", func(t *testing.T) {
		r1 := Rectangle{
			width:   10.0,
			breadth: 40.0,
		}
		got := r1.Perimeter()
		want := 100.0
		checkgotandwant(t, got, want)
	})
	t.Run("Testing the perimeter of circle", func(t *testing.T) {
		c1 := Circle{
			radius: 5.0,
		}
		got := c1.Perimeter()
		want := 31.42
		checkgotandwant(t, got, want)
	})
}
func TestArea(t *testing.T) {

	checkgotandwant := func(t testing.TB, got float64, want float64) {
		if got != want {
			t.Errorf("Expected: %.2f Received: %.2f", got, want)
		}
	}

	t.Run("Testing the area of rectangle", func(t *testing.T) {
		r2 := Rectangle{
			width:   5.0,
			breadth: 2.0,
		}
		got := r2.Area()
		want := 10.0
		checkgotandwant(t, got, want)
	})
	t.Run("Testing the Area of circle", func(t *testing.T) {
		c2 := Circle{
			radius: 10.0,
		}
		got := c2.Area()
		want := 314.16
		checkgotandwant(t, got, want)
	})
}
