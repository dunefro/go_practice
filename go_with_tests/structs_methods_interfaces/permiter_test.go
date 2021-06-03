package perimeter

import "testing"

type Shape interface {
	Perimeter() float64
	Area() float64
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

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("Expected: %g Received: %g", got, want)
		}
	}

	t.Run("Testing the area of rectangle", func(t *testing.T) {
		r2 := Rectangle{
			width:   5.0,
			breadth: 2.0,
		}

		want := 10.0
		checkArea(t, r2, want)
	})
	t.Run("Testing the Area of circle", func(t *testing.T) {
		c2 := Circle{
			radius: 10.0,
		}
		want := 314.16
		checkArea(t, c2, want)
	})
}
