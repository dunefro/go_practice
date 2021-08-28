package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	recordWin(name string)
}
type PlayerServer struct {
	store PlayerStore
}
type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}
type InMemoryPlayerStore struct {
	scores map[string]int
}

func (i *InMemoryPlayerStore) GetPlayerScore(player string) int {
	score := i.scores[player]
	return score
}

func (i *InMemoryPlayerStore) recordWin(name string) {
	i.scores[name]++
}

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5000", server))
}
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	w.WriteHeader(http.StatusAccepted)
	p.store.recordWin(player)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {

	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}

func (s *StubPlayerStore) GetPlayerScore(player string) int {
	score := s.scores[player]
	return score
}
func (s *StubPlayerStore) recordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}
