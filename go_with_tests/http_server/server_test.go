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
		nil,
	}
	server := &PlayerServer{&store}
	t.Run("Return Pepper's score", func(t *testing.T) {
		request := newgetscorerequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})
	t.Run("Return Floyd's score", func(t *testing.T) {
		request := newgetscorerequest("floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})
	t.Run("Player not present", func(t *testing.T) {
		request := newgetscorerequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
	}
	server := &PlayerServer{&store}
	t.Run("POST request", func(t *testing.T) {
		request := newPostWinRequest("Pepper")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusAccepted)
		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}
	})
}
func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := InMemoryPlayerStore{
		map[string]int{},
	}
	server := PlayerServer{&store}
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newgetscorerequest(player))
	assertStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "3")
}

func newPostWinRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return request
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
func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Didn't get the correct status. got %d want %d", got, want)
	}
}
