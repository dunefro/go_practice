package main

import (
	"fmt"
	"strings"
)

type Person struct {
	name    string
	age     int
	hobbies []string
}

func main() {
	p := &Person{
		name: "A",
		age:  12,
		hobbies: []string{
			"sports",
			"youtube",
			"music",
		},
	}
	fmt.Println(p.message())
	p.addHobbie("movies")
	fmt.Println(p.message())
	p = nil
	fmt.Println(p.message())

}

func (p *Person) message() string {
	if p == nil {
		return "Are you a person ?"
	}
	return fmt.Sprintf("My name is %s. I am %d years old. My hobbies are %s.", p.name, p.age, strings.Join(p.hobbies, ","))
}

func (p *Person) addHobbie(hobby string) {
	p.hobbies = append(p.hobbies, hobby)
}
