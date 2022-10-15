package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func MockWebsiteTester(s string) bool {
	if s == "google.com" {
		return false
	}
	return true
}

func SlowWebsiteTesterMock(s string) bool {
	time.Sleep(20 * time.Millisecond)
	return MockWebsiteTester(s)
}

var websites = []string{
	"amazon.com",
	"youtube.com",
	"google.com",
}

func TestCheckWebsites(t *testing.T) {
	t.Run("Check websites", func(t *testing.T) {

		got := CheckWebsites(MockWebsiteTester, websites)
		want := map[string]bool{
			"amazon.com":  true,
			"youtube.com": true,
			"google.com":  false,
		}

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v but want %v", got, want)
		}
	})

	t.Run("Test slow websites", func(t *testing.T) {

	})
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "some url"
	}

	/* The setup shouldn't be timed */
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CheckWebsites(SlowWebsiteTesterMock, urls)
	}
}
