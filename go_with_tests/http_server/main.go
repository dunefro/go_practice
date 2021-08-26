package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	handler := http.HandlerFunc(PlayerServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	if player == "Pepper" {
		fmt.Fprintf(w, "20")
	} else if player == "floyd" {
		fmt.Fprintf(w, "10")
	}
}