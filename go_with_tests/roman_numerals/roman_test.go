package main

import "testing"

func TestConvertToRoman(t *testing.T) {
	got := ConvertToRoman(1)
	want := "I"

	if got != want {
		t.Errorf("Recieved: %q, Expected: %q", got, want)
	}
}
