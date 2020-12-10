package maps

import "testing"

func TestDictionary(t *testing.T) {
	t.Run("Search for word", func(t *testing.T) {
		dictionary := Dictionary{"test": "definition of test"}

		got := dictionary.Search("test")
		want := "definition of test"

		assertStrings(t, got, want)

	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}
