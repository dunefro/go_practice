package main

import (
	"log"
)

type customErr struct {
	s string
}

func (c customErr) Error() string {
	return c.s
}

func main() {
	foo(customErr{"Some error"})
}
func foo(err error) {
	log.Println(err.Error())
}
