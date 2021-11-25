package main

import (
	"fmt"
	"math"
)

const (
	pi = math.Pi
)

type geometry interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	length  float64
	breadth float64
}

func (c circle) area() float64 {
	return pi * c.radius * c.radius
}

func (r rectangle) area() float64 {
	return r.length * r.breadth
}

func (c circle) perimeter() float64 {
	return 2 * pi * c.radius
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.length + r.breadth)
}

func calArea(g geometry) float64 {
	return g.area()
}
func calPerimeter(g geometry) float64 {
	return g.perimeter()
}

func main() {
	r1 := rectangle{4.0, 5.0}
	c1 := circle{3.0}
	fmt.Println(c1.area())
	fmt.Println(r1.area())
	fmt.Println(calArea(c1))
	fmt.Println(calArea(r1))
	fmt.Println(calPerimeter(c1))
	fmt.Println(calPerimeter(r1))
}
