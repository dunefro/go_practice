package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	p := Person{
		name: "Vedant Pareek",
		age:  23,
	}
	q := Person{
		name: "Harry Potter",
		age:  11,
	}
	r := &Person{
		name: "Albus Dumbledore",
		age:  150,
	}
	s := &Person{
		name: "Rubeius Hagrid",
		age:  90,
	}
	fmt.Println((&p).message()) // passing address
	fmt.Println(q.message())    // passing variable directly
	fmt.Println(r.message())    // passing pointer directly
	fmt.Println((*s).message()) // passing dereferenced pointer
	fmt.Println((&Person{
		name: "Ron Weasely",
		age:  10,
	}).message())
}

func (p *Person) message() string {
	return fmt.Sprintf("My name is %s. I am %d years old.", p.name, p.age)
}
