package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("Return Pepper's score", func(t *testing.T) {
		request := newgetscorerequest("Pepper")
		response := httptest.NewRecorder()

		PlayerServer(response, request)
		got := response.Body.String()
		want := "20"
		assertResponseBody(t, got, want)
	})
	t.Run("Return Floyd's score", func(t *testing.T) {
		request := newgetscorerequest("floyd")
		response := httptest.NewRecorder()

		PlayerServer(response, request)
		got := response.Body.String()
		want := "10"
		assertResponseBody(t, got, want)
		// if got != want {
		// 	t.Errorf("got %q want %q", got, want)
		// }
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
