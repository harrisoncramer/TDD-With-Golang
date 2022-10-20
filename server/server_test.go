package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPlayerServer(t *testing.T) {
	t.Run("Handler returns the player's score", func(t *testing.T) {

		want := "20"
		request, _ := http.NewRequest(http.MethodGet, "/players/me", nil)

		/* Returns a special response writer spy that records
				   the outuput for later inspection by tests,
		       on response.Body.String() */
		response := httptest.NewRecorder()

		/*
		 PlayerServer implements the Handler interface:
		 type Handler interface {
		   ServeHTTP (ResponseWriter, *Request)
		 }
		*/
		PlayerServer(response, request)

		got := response.Body.String()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Handler returns another player's score", func(t *testing.T) {
		want := "10"
		request, _ := http.NewRequest(http.MethodGet, "/players/you", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})
}
