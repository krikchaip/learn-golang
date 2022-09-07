package server

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

// implements: http.Handler
type PlayerServer struct {
	store PlayerStore
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	return &PlayerServer{store}
}

func (s *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := http.NewServeMux()

	router.HandleFunc("/league", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.HandleFunc("/players/", func(w http.ResponseWriter, r *http.Request) {
		player := strings.TrimPrefix(r.URL.Path, "/players/")

		switch r.Method {
		case http.MethodGet:
			s.showScore(w, player)
		case http.MethodPost:
			s.processWin(w, player)
		}
	})

	router.ServeHTTP(w, r)
}

func (s *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := s.store.GetPlayerScore(player)

	if score > 0 {
		fmt.Fprint(w, score)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func (s *PlayerServer) processWin(w http.ResponseWriter, player string) {
	s.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
