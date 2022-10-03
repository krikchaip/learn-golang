package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"21-acceptance-testing/lib/util"
	"22-building-application/entity"
)

// implements: http.Handler
type PlayerServer struct {
	store entity.PlayerStore
	game  entity.Game

	// ** interface embedding (like `implements IFoo` in other languages)
	http.Handler

	// // ?? alternative to interface embedding
	// router *http.ServeMux

	template *template.Template
}

func NewPlayerServer(
	store entity.PlayerStore,
	game entity.Game,
) *PlayerServer {
	// loading a HTML template for /game
	tmpl, err := loadTemplate()
	if err != nil {
		panic(err)
	}

	// s := &PlayerServer{store, router}
	s := &PlayerServer{
		store:    store,
		game:     game,
		template: tmpl,
	}

	// this also implements http.Handler
	router := http.NewServeMux()
	s.Handler = router

	router.HandleFunc("/league", s.leagueHandler)
	router.HandleFunc("/players/", s.playersHandler)
	router.HandleFunc("/game", s.gameHandler)
	router.HandleFunc("/ws", s.wsHandler)

	return s
}

func loadTemplate() (*template.Template, error) {
	_, filename, _, _ := runtime.Caller(0)
	root, _ := util.FindRoot(filepath.Dir(filename))
	filepath := filepath.Join(root, "view/game.html")

	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		return nil, fmt.Errorf("problem loading template %s", err.Error())
	}

	return tmpl, nil
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

func (s *PlayerServer) gameHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	s.template.Execute(w, nil)
}

func (s *PlayerServer) wsHandler(w http.ResponseWriter, r *http.Request) {
	ws := newPlayerServerWS(w, r)

	msg := ws.WaitForMsg()
	nPlayers, _ := strconv.Atoi(msg)

	// ?? io.Discard -> .Write() noop
	// s.game.Start(io.Discard, nPlayers)

	s.game.Start(ws, nPlayers)

	winner := ws.WaitForMsg()
	s.store.RecordWin(winner)
}
