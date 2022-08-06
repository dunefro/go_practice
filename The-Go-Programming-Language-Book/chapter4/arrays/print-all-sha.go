package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	shaType, word := "", ""
	if len(os.Args) <= 2 {
		shaType = "256"
		word = os.Args[1]
	} else {
		if os.Args[1] == "-x" {
			shaType = "256"
			word = os.Args[2]
		} else if os.Args[1] == "-y" {
			shaType = "512"
			word = os.Args[2]
		} else {
			panic("Only -x(SHA256) or -y(SHA512) accepted")
		}
	}
	fmt.Println(getSHA(shaType, word))
}
func getSHA(s, w string) string {
	if s == "256" {
		return fmt.Sprintf("%x", sha256.Sum256([]byte(w)))
	} else {
		return fmt.Sprintf("%x", sha512.Sum512([]byte(w)))
	}
}
