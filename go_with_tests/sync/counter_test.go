package main

import "testing"

func TestCounter(t *testing.T) {
	t.Run("Incrementing the counter to 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		want := 3
		assertCounter(t, counter, want)

	})
}
func assertCounter(t *testing.T, got Counter, want int) {

	t.Helper()

	if got.Value() != want {
		t.Errorf("Expected %d recieved %d", want, got.Value())
	}
}
