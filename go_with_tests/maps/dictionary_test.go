package maps

import "testing"

func TestDictionary(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("Known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})
	t.Run("Unknown Word", func(t *testing.T) {
		_, err := dictionary.Search("random")
		want := "Could not find the word you are looking for"
		if err == nil {
			t.Fatal("Expected error here")
		}
		assertStrings(t, err.Error(), want)
	})
}

func assertStrings(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Expected: %q Received: %q", want, got)
	}
}
