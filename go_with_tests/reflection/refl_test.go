package reflection

import "testing"

func TestWalk(t *testing.T) {

	expected := "Vedant"
	var got []string

	x := struct {
		Name string
	}{expected}

	walk(x, func(input string) {
		got = append(got, input)
	})

	if got[0] != expected {
		t.Errorf("Expected %q, Recieved %q", expected, got[0])
	}
}
