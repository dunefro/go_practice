package main

import "fmt"

type human interface {
	speak()
}

type person struct {
	first string
	last  string
}

func saySomething(h human) {
	h.speak()
}

func (p *person) speak() {
	fmt.Println(p.first, p.last)
}

func main() {
	p1 := person{
		first: "vedant",
		last:  "pareek",
	}
	saySomething(&p1)
}
