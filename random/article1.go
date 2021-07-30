package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	Peter := person{"Vedant", "Pareek"}
	fmt.Println(Peter.fullName())
}

func (p person) fullName() string {
	return p.firstName + " " + p.lastName
}
