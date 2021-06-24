package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("Incrementing the counter to 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		want := 3
		assertCounter(t, counter, want)

	})

	t.Run("Incrementing the counter to big value", func(t *testing.T) {

		wantedCount := 1000
		var wg sync.WaitGroup
		var mu sync.Mutex
		wg.Add(wantedCount)
		counter := Counter{}
		for i := 0; i < 1000; i++ {
			go func() {
				mu.Lock()
				counter.Inc()
				mu.Unlock()
				wg.Done()
			}()
		}
		wg.Wait()
		assertCounter(t, counter, wantedCount)

	})
}
func assertCounter(t *testing.T, got Counter, want int) {

	t.Helper()

	if got.Value() != want {
		t.Errorf("Expected %d recieved %d", want, got.Value())
	}
}
