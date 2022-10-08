package iteration

import (
	"reflect"
	"testing"
)

func assertTestPassing(got string, want string, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertTestPassingSlice(got []string, want []string, t *testing.T) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}

// Returns a string that repeats the first argument n times
func TestIterate(t *testing.T) {
	got := Iterate("a", 10)
	want := "aaaaaaaaaa"
	assertTestPassing(got, want, t)
}

func TestUppercase(t *testing.T) {
	got := Uppercase("harry")
	want := "Harry"
	assertTestPassing(got, want, t)
}

func TestWords(t *testing.T) {
	t.Run("A normal sentence", func(t *testing.T) {
		got := Words("This is a sentence")
		want := []string{"This", "is", "a", "sentence"}
		assertTestPassingSlice(got, want, t)
	})

	t.Run("A sentence with numbers", func(t *testing.T) {
		got := Words("Wow 100 is a lot")
		want := []string{"Wow", "100", "is", "a", "lot"}
		assertTestPassingSlice(got, want, t)
	})

	t.Run("A blank string", func(t *testing.T) {
		got := Words("")
		want := []string{}
		assertTestPassingSlice(got, want, t)
	})

	t.Run("Only whitespace", func(t *testing.T) {
		got := Words("")
		want := []string{}
		assertTestPassingSlice(got, want, t)
	})

	t.Run("No spaces at all", func(t *testing.T) {
		got := Words("thisisnotasentence")
		want := []string{"thisisnotasentence"}
		assertTestPassingSlice(got, want, t)
	})
}

/*
	We can benchmark the performance of an operation, Golang determines

how many times to run the function to return a good benchmark.
We can run them with: go test -bench=.
*/
func BenchmarkIterate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iterate("q", 10)
	}
}
