package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Player struct {
	Name string
	Wins int
}

type PlayerStore interface {
	GetPlayerScore(name string) int
	GetLeagueTable() []Player
	RecordWin(name string)
}

// implements: http.Handler
type PlayerServer struct {
	store PlayerStore

	// ** interface embedding (like `implements IFoo` in other languages)
	http.Handler

	// // ?? alternative to interface embedding
	// router *http.ServeMux
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	// this also implements http.Handler
	router := http.NewServeMux()

	// s := &PlayerServer{store, router}
	s := &PlayerServer{store: store}
	s.Handler = router

	router.HandleFunc("/league", s.leagueHandler)
	router.HandleFunc("/players/", s.playersHandler)

	return s
}

// // ?? alternative to interface embedding
// func (s *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	s.router.ServeHTTP(w, r)
// }

func (s *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	// ** transform a struct into a JSON string and then Write()
	json.NewEncoder(w).Encode(s.store.GetLeagueTable())

	// // ?? alternative to json.Encoder (using json.Marshal)
	// bytes, _ := json.Marshal(s.store.GetLeagueTable())
	// w.Write(bytes)
}

func (s *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodGet:
		s.showScore(w, player)
	case http.MethodPost:
		s.processWin(w, player)
	}
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
