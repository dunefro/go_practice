package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func main() {
	p1 := []person{{
		Name: "Vedant Pareek",
		Age:  22,
	}}
	v, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("Not able to marshal")
	}
	fmt.Println(string(v))
}
