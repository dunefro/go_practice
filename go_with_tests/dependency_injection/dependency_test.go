package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Vedant")

	got := buffer.String()
	want := "Hello Vedant"
	if got != want {
		t.Errorf("Expected: %q Received: %q", want, got)
	}
}
