package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func assertResponseBody(t *testing.T, got string, want string) {
	if got != want {
		t.Fatalf("got %q but want %q", got, want)
	}
}

func TestGETPlayers(t *testing.T) {

	server := &PlayerServer{
		Store: &StubPlayerStore{
			scores: map[string]int{
				"peter": 10,
				"harry": 20,
			},
		},
	}

	t.Run("Should return Peter's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/peter", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertResponseBody(t, response.Body.String(), "10")
	})
	t.Run("Should return Harry's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/harry", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertResponseBody(t, response.Body.String(), "20")
	})
}
