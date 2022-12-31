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

func assertResponseStatus(t *testing.T, got int, want int) {
	if got != want {
		t.Fatalf("got status %d but want %d", got, want)
	}
}

func TestGETPlayers(t *testing.T) {

	server := &PlayerServer{
		Store: &StubPlayerStore{
			scores: map[string]int{
				"peter": 10,
				"ruby":  20,
			},
		},
	}

	t.Run("Should return Peter's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/peter", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertResponseBody(t, response.Body.String(), "10")

		gotStatus := response.Result().StatusCode
		assertResponseStatus(t, gotStatus, 200)
	})

	t.Run("Should return Harry's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/ruby", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertResponseBody(t, response.Body.String(), "20")

		gotStatus := response.Result().StatusCode
		assertResponseStatus(t, gotStatus, 200)
	})

	t.Run("Should 404 if user does not exist", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/will", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "Player not found")

		gotStatus := response.Result().StatusCode
		assertResponseStatus(t, gotStatus, 404)
	})

	t.Run("Should return 200 on POST", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/players/peter", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		gotStatus := response.Result().StatusCode
		assertResponseStatus(t, gotStatus, 200)

	})
}
