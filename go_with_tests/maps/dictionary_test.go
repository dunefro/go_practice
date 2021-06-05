package maps

import "testing"

func TestDictionary(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test"}
	got := Search(dictionary, "test")
	want := "this is just a test"

	if got != want {
		t.Errorf("Expected: %q Received: %q", want, got)
	}
}
