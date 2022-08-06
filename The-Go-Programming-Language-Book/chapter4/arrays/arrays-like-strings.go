package main

import "fmt"

type currency int

const (
	USD currency = iota
	INR
	GBP
	KRW
)

func main() {
	// [...] -> to calculate the length of the array
	// with this we can use it as a simple map structure if the value doesn't change much
	symbol := [...]string{USD: "$", INR: "₹", GBP: "£", KRW: "₩"}
	fmt.Println(symbol[USD])
}
