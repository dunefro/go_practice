package geometry

import "math"

type Point struct {
	X float64
	Y float64
}
type Path []Point

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, p.Y-q.Y)
}

func (p Path) Distance() float64 {
	var totalDistance float64
	for i, q := range p {
		if i == len(p)-1 {
			totalDistance += q.Distance(p[0])
		}
		totalDistance += q.Distance(p[i+1])
	}
	return totalDistance
}
