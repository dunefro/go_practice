package main

import "testing"

func TestMysum(t *testing.T) {
	s := mySum(2, 3)
	expected := 6
	if s != expected {
		t.Error("Expected", expected, "Got", s)
	}
}
