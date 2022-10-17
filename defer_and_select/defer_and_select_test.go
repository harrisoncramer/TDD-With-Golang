package deferandselect

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

/* When testing servers we can wrap the handler in httptest.NewServer */
func makeServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRacer(t *testing.T) {
	t.Run("Should return the URL of the faster server", func(t *testing.T) {
		slowServer := makeServer(20 * time.Millisecond)
		fastServer := makeServer(0)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := Racer(slowURL, fastURL, 10000)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Returns an error on 10 second timeout", func(t *testing.T) {
		superSlowServer := makeServer(30 * time.Millisecond)
		otherSlowServer := makeServer(40 * time.Millisecond)
		defer superSlowServer.Close()
		defer otherSlowServer.Close()

		slowURL := superSlowServer.URL
		slowerURL := otherSlowServer.URL

		_, err := Racer(slowURL, slowerURL, 10)
		if err == nil {
			t.Errorf("Expected timeout error but didn't get one")
		}
	})
}
