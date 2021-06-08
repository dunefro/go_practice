package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	sleeper := Spysleeper{
		Calls: 3,
	}
	Countdown(buffer, &sleeper)
	got := buffer.String()
	want := `3
2
1
Go!`
	if got != want {
		t.Errorf("Expected: %q Received: %q", got, want)
	}
}
