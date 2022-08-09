package main

import "os"

func main() {
	f, err := os.Create("./newfile.txt")
	if err != nil {
		panic(err)
	}
	b := []byte("Hello world")
	f.Write(b)

}
