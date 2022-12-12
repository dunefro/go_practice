package main

import "fmt"

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p := &Person{"Vedant", "Pareek", 23}
	fmt.Println(p)
}

func (p *Person) String() string {
	return fmt.Sprintf("Name: %s %s, Age: %d", p.First, p.Last, p.Age)
}
