package main

import (
	"fmt"
	"log"
	"net/http"
)

var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Println("Listening ...")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	count++
	log.Println(r.URL.Path)
	fmt.Fprintf(w, "URL: %s", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	count++
	fmt.Fprintf(w, "Counter: %d", count)
}
