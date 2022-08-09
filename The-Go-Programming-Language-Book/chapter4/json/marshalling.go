package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "Vedant",
		Last:  "Pareek",
		Age:   23,
	}
	p2 := person{
		First: "Harry",
		Last:  "Potter",
		Age:   11,
	}
	p3 := person{
		First: "A",
		Last:  "B",
	}
	var people = []person{p1, p2, p3}
	fmt.Println(people)
	m, err := json.Marshal(people)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(m))
	n, err := json.MarshalIndent(people, "", " ")
	fmt.Println(string(n))
}
