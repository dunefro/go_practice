package main

import (
	"fmt"
	"method-geometry/geometry"
)

func main() {
	P := geometry.Point{
		X: 7.0,
		Y: 1.0,
	}
	Q := geometry.Point{
		X: 3.0,
		Y: 4.0,
	}
	fmt.Println(P.Distance(Q))
	// perimeter of a right angle triangle
	path := geometry.Path{
		geometry.Point{1, 1},
		geometry.Point{5, 1},
		geometry.Point{5, 4},
	}
	fmt.Println(path.Distance())

}
