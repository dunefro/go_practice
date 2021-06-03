package perimeter

import "math"

func roundoff(f float64) float64 {
	return math.Round(f*100) / 100
}

func RectanglePermiter(rectangle Rectangle) float64 {
	return roundoff(2 * (rectangle.width + rectangle.breadth))
}

func RectangleArea(rectangle Rectangle) float64 {
	return roundoff(rectangle.width * rectangle.breadth)
}

func CirclePerimeter(circle Circle) float64 {
	return roundoff(2 * math.Pi * circle.radius)
}
func CircleArea(circle Circle) float64 {
	return roundoff(math.Pi * circle.radius * circle.radius)
}
