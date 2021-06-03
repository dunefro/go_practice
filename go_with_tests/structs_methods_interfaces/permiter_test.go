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

func TestPerimeter(t *testing.T) {

	checkPerimeter := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("Expected: %g Received: %g", got, want)
		}
	}
	perimeterTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Rectangle", Rectangle{10.0, 40.0}, 100.0},
		{"Circle", Circle{5.0}, 31.42},
	}

	for i := 0; i < len(perimeterTests); i++ {
		t.Run(perimeterTests[i].name, func(t *testing.T) {
			checkPerimeter(t, perimeterTests[i].shape, perimeterTests[i].want)
		})
	}
}
func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("Expected: %g Received: %g", got, want)
		}
	}
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{5.0, 2.0}, 10.0},
		{Circle{10.0}, 314.16},
	}
	for i := 0; i < len(areaTests); i++ {
		checkArea(t, areaTests[i].shape, areaTests[i].want)
	}
}
