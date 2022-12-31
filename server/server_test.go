package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

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
		assertResponseStatus(t, gotStatus, http.StatusNotFound)
	})

	t.Run("Should return 200 on POST", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/players/peter", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		gotStatus := response.Result().StatusCode
		assertResponseStatus(t, gotStatus, http.StatusOK)

	})

	t.Run("Should record win on a POST", func(t *testing.T) {
		playerName := "sam"
		store := StubPlayerStore{
			map[string]int{},
			nil,
		}
		server = &PlayerServer{
			Store: &store,
		}

		request, _ := http.NewRequest(http.MethodPost, "/players/sam", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertResponseStatus(t, response.Result().StatusCode, http.StatusOK)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d wincalls but wanted %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != playerName {
			t.Errorf("got %s win but wanted %s win", store.winCalls[0], playerName)
		}

	})
}

/* Assertion helpers */
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
func assertEqual(t *testing.T, got any, want any) {
	if got != want {
		t.Fatalf("got %v but want %v", got, want)
	}
}
