package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(w io.Writer) {
	for i := 3; i >= 1; i-- {
		fmt.Fprint(w, i, "\n")
	}
	fmt.Fprint(w, "Go!")
}

func main() {
	Countdown(os.Stdout)
}
