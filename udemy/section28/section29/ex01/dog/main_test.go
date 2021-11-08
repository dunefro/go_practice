package dog

import "testing"

func TestYears(t *testing.T) {
	dogAge := 5
	got := Years(dogAge)
	expected := 35
	if got != expected {
		t.Error("Expected", expected, "Got", got)
	}
}
