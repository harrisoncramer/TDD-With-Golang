package mocking

import (
	"bytes"
	"reflect"
	"testing"
)

/*
We can implement the same interface that a regular sleeper would use
(in this case the time.Sleep() method) and inject the dependency to spy
on its behavior.
*/
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

/*
In a more complex example, we can add some functionality to the sleep method
and the write method to keep track of the order of operations for the injected
dependency.
*/
type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

func TestCountdown(t *testing.T) {
	t.Run("Prints expected output", func(t *testing.T) {
		buffer := bytes.Buffer{}
		spySleeper := SpySleeper{}
		Countdown(&buffer, &spySleeper)

		got := buffer.String()
		want := `3
2
1
go!`
		assertTestPassing(t, got, want)

		if spySleeper.Calls != 3 {
			t.Errorf("expected %d calls, got %d calls", 3, spySleeper.Calls)
		}
	})

	t.Run("Sleeps before every write", func(t *testing.T) {
		spySleepPrinter := SpyCountdownOperations{}
		Countdown(&spySleepPrinter, &spySleepPrinter)

		want := []string{
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("got calls %v wanted %v", spySleepPrinter.Calls, want)
		}
	})
}

func assertTestPassing(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
