package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout // OK
	fmt.Println(w)
	w = new(bytes.Buffer) // OK
	fmt.Println(w)
}
