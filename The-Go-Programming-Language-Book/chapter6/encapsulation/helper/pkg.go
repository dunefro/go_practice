package helper

import "fmt"

// struct type is public
type People struct {
	names   []string // private and encapsulated
	Country string
}

func Group() *People {
	return &People{
		names:   []string{"Harry", "Ron", "Dumbledore"},
		Country: "England",
	}
}

func (p *People) PrintPeople() {
	for _, people := range p.names {
		fmt.Println(people)
	}
}
