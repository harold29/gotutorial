package maps

import "testing"

func TestDictionary(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("with known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("with unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		want := "could not find the word you were looking for"

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertStrings(t, err.Error(), want)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}
