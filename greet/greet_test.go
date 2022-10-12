package greet

import (
	"bytes"
	"testing"
)

/*
This testing approach lets us inject a dependency (in this case) a buffer,
which implements the io.Writer interface) in order to make our Greet function
testable. We could then use greet with standard output, a file, etc.
*/
func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Harry")
	got := buffer.String()
	want := "Hello, Harry"
	assertTestPassing(t, got, want)
}

func assertTestPassing(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
