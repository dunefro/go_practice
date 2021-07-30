package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	Person := person{"Vedant", "Pareek"}
	fmt.Println(Person.fullName())
	fmt.Println("After changing the firstName")
	// Person.changeFirstName("Mr. Vedant")
	// fmt.Println(Person.fullName())
	PointerPerson := &Person
	PointerPerson.changeFirstName("Mr Vedant")
	fmt.Println(Person.fullName())
}

func (p person) fullName() string {
	return p.firstName + " " + p.lastName
}

// func (p person) changeFirstName(s string) {
// 	p.firstName = s
// }

func (p *person) changeFirstName(s string) {
	(*p).firstName = s
}
