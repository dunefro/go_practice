package maps

import "testing"

func TestDictionary(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	got := Search(dictionary, "test")
	want := "this is just a test"
	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Expected: %q Received: %q", want, got)
	}
}
