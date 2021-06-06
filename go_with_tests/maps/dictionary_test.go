package maps

import "testing"

func TestDictionary(t *testing.T) {
	text := "this is just a test"
	dictionary := Dictionary{"test": text}

	t.Run("Known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := text
		assertStrings(t, got, want)
	})
	t.Run("Unknown Word", func(t *testing.T) {
		_, err := dictionary.Search("random")
		if err == nil {
			t.Fatal("Expected error here")
		}
		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	text := "this is just a test"

	t.Run("Adding non-existent key", func(t *testing.T) {
		dictionary := Dictionary{}
		//Adding the key here
		errAdd := dictionary.Add("test", text)
		// Searching for the added key
		got, errSearch := dictionary.Search("test")
		want := text
		if errAdd != nil {
			t.Fatal("Was not expectng any error here but recieved one")
		}
		if errSearch != nil {
			t.Fatal("Was not expecting an error here")
		}
		assertStrings(t, got, want)
	})
	t.Run("Adding existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": text}
		errAdd := dictionary.Add("test", "some text that should not get overridden")
		got, errSearch := dictionary.Search("test")
		if errAdd == nil {
			t.Error("Was expecting an error herer")
		}
		if errSearch != nil {
			t.Fatal("Was not expecting the error here")
		}
		want := text
		assertStrings(t, got, want)
	})

}

func TestUpdate(t *testing.T) {
	text := "this is just a test"
	dictionary := Dictionary{"test": text}
	newText := "something new"
	dictionary.Update("test", newText)
	got, errSearch := dictionary.Search("test")
	if errSearch != nil {
		t.Fatal("Expected the error not to come")
	}
	assertStrings(t, got, newText)
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
