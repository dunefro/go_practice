package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Age   int    `json:"age,omitempty"`
}

func main() {
	people := `[{"First":"Vedant","Last":"Pareek","Age":23},{"First":"Harry","Last":"Potter"}]`
	var persons []person
	json.Unmarshal([]byte(people), &persons)
	for _, v := range persons {
		fmt.Println(v)
	}
}
