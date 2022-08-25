package main

import (
	"encap/helper"
	"fmt"
)

func main() {
	g := helper.Group()
	g.PrintPeople() // public
	// fmt.Println(g.names) // gives error as this is encapsulated
	fmt.Println(g.Country) // public
}
