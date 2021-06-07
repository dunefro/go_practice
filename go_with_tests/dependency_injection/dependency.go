package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func Greet(writer io.Writer, name string) {
	name = strings.Trim(name, "/")
	fmt.Fprintf(writer, "Hello %s", name)
}

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, r.URL.Path)
	// if r.URL.Path != "/" {
	// 	Greet(w, "Vedant")
	// }
	// if r.URL.Path != "/world"
}

func main() {
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreetHandler)))
}
