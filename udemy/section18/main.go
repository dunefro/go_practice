package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Our password
	s := "Password123"
	cost := 10
	encrypt_password, err := bcrypt.GenerateFromPassword([]byte(s), cost)
	fmt.Println(encrypt_password)
	fmt.Println(err)
}
