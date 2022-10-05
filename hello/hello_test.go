package hello

import "testing"

func assertTestPassing(t *testing.T, got string, want string) {
	t.Helper() // Keeps line numbers correct
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
func TestHello(t *testing.T) {

	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertTestPassing(t, got, want)
	})

	t.Run("Saying 'Hello World' with an empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertTestPassing(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Peter", "Spanish")
		want := "Hola, Peter"
		assertTestPassing(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Peter", "French")
		want := "Bonjour, Peter"
		assertTestPassing(t, got, want)
	})
}
