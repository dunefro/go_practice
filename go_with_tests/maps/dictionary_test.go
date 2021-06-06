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
		if err == nil {
			t.Fatal("Expected error here")
		}
		assertError(t, err, errorNotFound)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	//Adding the key here
	dictionary.Add("test", "this is just a test")
	// Searching for the added key
	got, err := dictionary.Search("test")
	want := "this is just a test"
	if err != nil {
		t.Fatal("Was not expecting an error here")
	}
	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Expected: %q Received: %q", want, got)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
