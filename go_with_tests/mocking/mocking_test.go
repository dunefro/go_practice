package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	Countdown(buffer)
	got := buffer.String()
	want := "3"
	if got != want {
		t.Errorf("Expected: %q Received: %q", got, want)
	}
}
