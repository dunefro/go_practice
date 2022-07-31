package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "skips newline character")
var sep = flag.String("s", " ", "join arguments with this")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
