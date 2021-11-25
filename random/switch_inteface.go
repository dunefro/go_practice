package main

import "fmt"

type wizard interface {
	intro() string
}

type men struct {
	name  string
	age   int
	house string
}

type women struct {
	name  string
	age   int
	house string
}

func (m men) intro() string {
	return fmt.Sprintln("I am", m.name, "of age", m.age, "in house", m.house)
}

func (w women) intro() string {
	return fmt.Sprintln("I am", w.name, "of age", w.age, "in house", w.house)
}

func speak(wi wizard) string {
	return wi.intro()
}

func checkDormitory(wi wizard) string {
	var dormitory string
	switch wi.(type) {
	case men:
		dormitory = "boys"
	case women:
		dormitory = "girls"
	}
	return dormitory
}

func main() {
	m1 := men{"Vedant Pareek", 22, "Slytherin"}
	w1 := women{"Ginny Weasely", 18, "Gryffindor"}
	fmt.Println(speak(m1), "Your dormitory is", checkDormitory(m1))
	fmt.Println(speak(w1), "Your dormitory is", checkDormitory(w1))
	// fmt.Println("Your dormitory is", checkDormitory(m1))
	// fmt.Println("Your dormitory is", checkDormitory(w1))

}
