package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(PlayerServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "20")
}
