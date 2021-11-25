package main

import "fmt"

type name struct {
	first string
	last  string
}
type person struct {
	fullname name
	age      int
	id       int
}

func main() {
	p1 := person{
		fullname: name{
			first: "Vedant",
			last:  "Pareek",
		},
		age: 22,
		id:  7,
	}

	p1.printName()
}
func (p person) printName() {
	fmt.Println(p.fullname.first, p.fullname.last)
}
