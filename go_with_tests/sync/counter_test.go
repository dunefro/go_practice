package main

import "testing"

func TestCounter(t *testing.T) {
	t.Run("Incrementing the counter to 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()
		got := counter.Value()
		want := 3

		if got != want {
			t.Errorf("Expected %d recieved %d", want, got)
		}
	})
}
