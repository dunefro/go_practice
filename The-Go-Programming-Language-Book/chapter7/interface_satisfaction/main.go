package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	fmt.Println(w)
	w = new(bytes.Buffer)
	fmt.Println(w)
}
