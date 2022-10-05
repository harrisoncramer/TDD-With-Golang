package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	got := Add(2, 2)
	want := 4
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

/* Examples in test suites run as part of the suite
and can serve as documentation for the tested function */

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
