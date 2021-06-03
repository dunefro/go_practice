package perimeter

import "math"

func roundoff(f float64) float64 {
	return math.Round(f*100) / 100
}

func (r Rectangle) Perimeter() float64 {
	return roundoff(2 * (r.width + r.breadth))
}

func RectangleArea(rectangle Rectangle) float64 {
	return roundoff(rectangle.width * rectangle.breadth)
}

func (c Circle) Perimeter() float64 {
	return roundoff(2 * math.Pi * c.radius)
}
func CircleArea(circle Circle) float64 {
	return roundoff(math.Pi * circle.radius * circle.radius)
}
