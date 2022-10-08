package dictionary

import "testing"

func TestSearch(t *testing.T) {

	t.Run("Should find valid word", func(t *testing.T) {
		dictionary := Dictionary{"test": "This is a test"}
		got, err := dictionary.Search("test")
		want := "This is a test"
		assertNoError(t, err)
		assertTestPassing(t, got, want)
	})

	t.Run("Should error if word not found", func(t *testing.T) {
		dictionary := Dictionary{}
		got, err := dictionary.Search("some-word")
		assertTestPassing(t, got, "")
		assertError(t, err, ErrWordNotFound)
	})
}

func TestDefine(t *testing.T) {
	t.Run("Should add a new word to the dictionary", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Define("ice", "frozen water")
		got, err := dictionary.Search("ice")

		assertNoError(t, err)
		assertTestPassing(t, got, "frozen water")
	})

	t.Run("Should throw if the map is nil", func(t *testing.T) {
		dictionary := Dictionary(nil)
		err := dictionary.Define("ice", "frozen water")
		assertError(t, err, ErrDictUndefined)
	})

	t.Run("Should throw if the word is already defined", func(t *testing.T) {
		dictionary := Dictionary{"test": "This is a test"}
		err := dictionary.Define("test", "some new value")
		assertError(t, err, ErrWordExists("test", "This is a test"))
	})
}

func assertTestPassing(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t *testing.T, err error, want error) {
	t.Helper()
	if err == nil {
		t.Fatal("wanted err but didn't get one")
	}

	if err.Error() != want.Error() {
		t.Errorf("got error %q want error %q", err, want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("err with message %s", err.Error())
	}
}
