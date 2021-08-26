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
	fmt.Fprint(w, getplayerscore(player))
}
func getplayerscore(player string) (value string) {
	if player == "Pepper" {
		value = "20"
	} else if player == "floyd" {
		value = "10"
	}
	return
}
