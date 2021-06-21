package sel

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	slowServer := returnfakeserver(20 * time.Millisecond)
	fastServer := returnfakeserver(0 * time.Millisecond)

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := Racer(slowURL, fastURL)
	if want != got {
		t.Errorf("We expected %q to be fast but %q was faster", want, got)
	}
	slowServer.Close()
	fastServer.Close()
}

func returnfakeserver(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
