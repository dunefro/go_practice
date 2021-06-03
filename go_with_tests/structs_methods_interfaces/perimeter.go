package perimeter

import "math"

func roundoff(f float64) float64 {
	return math.Round(f*100) / 100
}

func (r Rectangle) Perimeter() float64 {
	return roundoff(2 * (r.width + r.breadth))
}

func (r Rectangle) Area() float64 {
	return roundoff(r.width * r.breadth)
}

func (c Circle) Perimeter() float64 {
	return roundoff(2 * math.Pi * c.radius)
}
func (c Circle) Area() float64 {
	return roundoff(math.Pi * c.radius * c.radius)
}
