package mocking

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep(time int) {
	s.Calls++
}

func TestCountdown(t *testing.T) {
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
}

func assertTestPassing(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
