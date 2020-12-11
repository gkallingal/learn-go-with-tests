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

		assertDefinition(t, dictionary, word, definition)

	})
	t.Run("Add existing word", func(t *testing.T) {
		word := "test"
		definition := "definition of test"
		err := dictionary.Add(word, definition)
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("Update existing word", func(t *testing.T) {
		word := "test"
		newDefinition := "new definition of test"
		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})
	t.Run("Update unknown word", func(t *testing.T) {
		word := "fish"
		definition := "definition of fish"
		err := dictionary.Update(word, definition)
		assertError(t, err, ErrCannotUpdate)
	})
	t.Run("Delete a word", func(t *testing.T) {
		word := "test"
		dictionary.Delete(word)

		_, err := dictionary.Search(word)
		assertError(t, err, ErrNotFound)
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
	if err == nil {
		if want == nil {
			return
		}
		t.Fatalf("expected error, got none")
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatalf("%q should have been added: %q", word, err)
	}
	assertStrings(t, got, definition)
}
