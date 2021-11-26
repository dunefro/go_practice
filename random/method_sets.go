package main

import "fmt"

type wizard interface {
	magic()
}

type person struct {
	name   string
	age    int
	wizard bool
}

func (p *person) printAge() {
	fmt.Println(p.age)
}

func (p person) printName() {
	fmt.Println(p.name)
}
func (p *person) magic() {
	if p.wizard {
		fmt.Println("You are allowed to magic")
	}
}
func check_magic(w wizard) {
	w.magic()
}

func main() {
	p1 := person{"Vedant Pareek", 22, false}
	p1.printAge()
	p1.printName()
	p2 := &person{"Harry Potter", 11, true}
	(*p2).printAge()
	(*p2).printName()
	// check_magic(p1)
	check_magic(p2)
}
