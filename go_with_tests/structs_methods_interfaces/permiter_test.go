package perimeter

import "testing"

type Shape interface {
	Perimeter() float64
}

type Rectangle struct {
	width   float64
	breadth float64
}

type Circle struct {
	radius float64
}

func TestPermiter(t *testing.T) {

	checkPerimeter := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("Expected: %g Received: %g", got, want)
		}
	}

	t.Run("Testing the perimeter of rectangle", func(t *testing.T) {
		r1 := Rectangle{
			width:   10.0,
			breadth: 40.0,
		}
		want := 100.0
		checkPerimeter(t, r1, want)
	})
	t.Run("Testing the perimeter of circle", func(t *testing.T) {
		c1 := Circle{
			radius: 5.0,
		}
		want := 31.42
		checkPerimeter(t, c1, want)
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
