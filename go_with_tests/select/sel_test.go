package sel

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("Fake server testing", func(t *testing.T) {
		slowServer := returnfakeserver(20 * time.Millisecond)
		fastServer := returnfakeserver(0 * time.Millisecond)

		// These statements are invoked here but executed once TestRacer completes
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := Racer(slowURL, fastURL)
		if want != got {
			t.Errorf("We expected %q to be fast but %q was faster", want, got)
		}
	})
	t.Run("Time more than 10 seconds", func(t *testing.T) {
		slowServer := returnfakeserver(20 * time.Millisecond)
		fastServer := returnfakeserver(30 * time.Millisecond)

		// These statements are invoked here but executed once TestRacer completes
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		// want := fastURL
		_, err := Racer(slowURL, fastURL)
		if err == nil {
			t.Errorf("We expected an error here")
		}
	})
}

func returnfakeserver(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
