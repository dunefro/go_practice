package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"floyd":  10,
		},
	}
	server := &PlayerServer{&store}
	t.Run("Return Pepper's score", func(t *testing.T) {
		request := newgetscorerequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertResponseBody(t, response.Body.String(), "20")
	})
	t.Run("Return Floyd's score", func(t *testing.T) {
		request := newgetscorerequest("floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertResponseBody(t, response.Body.String(), "10")
	})
}

func newgetscorerequest(player string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
	return request
}
func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
