package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	personsInJSON := `[{"name":"Vedant Pareek","age":22},{"name":"Harry Potter","age":11}]`
	bytesOfPerson := []byte(personsInJSON)
	var people []person
	err := json.Unmarshal(bytesOfPerson, &people)
	if err != nil {
		fmt.Println("Some error occurred")
	}
	for _, v := range people {
		fmt.Println("Name:", v.fullname(), "and age is:", v.getAge())
	}
}
func (p person) fullname() string {
	return p.Name
}
func (p person) getAge() int {
	return p.Age
}
