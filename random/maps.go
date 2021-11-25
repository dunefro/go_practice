package main

import "fmt"

func main() {
	m := map[string]int{
		"k1": 1,
		"k2": 2,
		"k3": 3,
	}
	fmt.Println(m)
	if v, ok := m["k4"]; ok {
		fmt.Println("key k4 exists and has a value", v)
	} else {
		fmt.Println("key k4 doesn't exist")
	}
	map_of_maps()
}
func map_of_maps() {
	m := map[string]map[string]string{
		"bond_james": {
			"hello": "hi",
		},
	}

	fmt.Println(m)
}
