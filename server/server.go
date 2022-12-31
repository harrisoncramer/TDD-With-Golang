package server

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) (int, error)
}

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) (int, error) {
	score, ok := s.scores[name]
	if !ok {
		return 0, errors.New("Player not found")
	}

	return score, nil
}

type PlayerServer struct {
	Store PlayerStore
}

func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	score, err := p.Store.GetPlayerScore(player)
	if err != nil {
		w.WriteHeader(404)
		io.WriteString(w, err.Error())
		return
	}
	fmt.Fprint(w, score)
}
func (p *PlayerServer) updateScore(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		p.showScore(w, r)
	case http.MethodPost:
		p.updateScore(w, r)
	}
}
