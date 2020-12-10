package maps

import (
	"testing"
)

func TestDictionary(t *testing.T) {
	dictionary := Dictionary{"test": "definition of test"}

	t.Run("Search for word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "definition of test"

		assertStrings(t, got, want)

	})
	t.Run("Search for unknown word", func(t *testing.T) {
		word := "beta"
		_, err := dictionary.Search(word)
		if err == nil {
			t.Fatal("Expected error, got none")
		}
		assertStrings(t, err.Error(), ErrNotFound.Error())
		assertError(t, err, ErrNotFound)

	})
	t.Run("Add new word", func(t *testing.T) {
		word := "delta"
		definition := "definition of delta"
		dictionary.Add(word, definition)

		got, err := dictionary.Search(word)
		if err != nil {
			t.Fatalf("%q should have been added", word)
		}
		assertStrings(t, got, definition)

	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}

func assertError(t *testing.T, err, want error) {
	t.Helper()
	if err != want {
		t.Errorf("Got %q, want %q", err, want)
	}
}
