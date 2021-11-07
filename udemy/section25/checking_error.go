package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Opening a file, if file exists it will not give error
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// Creating slice of bytes to read data from the file
	bs := make([]byte, 100)
	_, err = file.Read(bs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bs))
}
