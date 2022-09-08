package server_test

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"22-building-application/entity"
	server "22-building-application/server"
	util "22-building-application/util/testing"
)

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{scores: map[string]int{
		"Pepper": 20,
		"Floyd":  10,
	}}
	sv := server.NewPlayerServer(&store)

	t.Run("returns Pepper's score", func(t *testing.T) {
		req := util.NewScoreRequest("Pepper")
		res := httptest.NewRecorder()

		sv.ServeHTTP(res, req)

		util.AssertStatus(t, res.Code, http.StatusOK)
		util.AssertResponseBody(t, res.Body, "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		req := util.NewScoreRequest("Floyd")
		res := httptest.NewRecorder()

		sv.ServeHTTP(res, req)

		util.AssertStatus(t, res.Code, http.StatusOK)
		util.AssertResponseBody(t, res.Body, "10")
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		req := util.NewScoreRequest("Apollo")
		res := httptest.NewRecorder()

		sv.ServeHTTP(res, req)

		util.AssertStatus(t, res.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{scores: map[string]int{}}
	sv := server.NewPlayerServer(&store)

	t.Run("it returns accepted on POST", func(t *testing.T) {
		req := util.NewPostWinRequest("Pepper")
		res := httptest.NewRecorder()

		sv.ServeHTTP(res, req)

		util.AssertStatus(t, res.Code, http.StatusAccepted)

		got := store.winCalls
		want := []string{"Pepper"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v calls to RecordWin want %v", got, want)
		}
	})
}

func TestLeague(t *testing.T) {
	t.Run("it returns 200 on /league", func(t *testing.T) {
		store := StubPlayerStore{}
		sv := server.NewPlayerServer(&store)

		req := util.NewLeagueRequest()
		res := httptest.NewRecorder()

		sv.ServeHTTP(res, req)

		util.AssertStatus(t, res.Code, http.StatusOK)
	})

	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := entity.League{
			{Name: "Cleo", Wins: 32},
			{Name: "Chris", Wins: 20},
			{Name: "Tiest", Wins: 14},
		}

		store := StubPlayerStore{league: wantedLeague}
		sv := server.NewPlayerServer(&store)

		req := util.NewLeagueRequest()
		res := httptest.NewRecorder()

		sv.ServeHTTP(res, req)

		got := util.ParseLeagueFromResponse(t, res.Body)
		util.AssertStatus(t, res.Code, http.StatusOK)
		util.AssertLeagueTable(t, got, wantedLeague)
		util.AssertContentJSON(t, res.Result().Header)
	})
}

// implements: server.PlayerStore
type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   entity.League
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) GetLeagueTable() entity.League {
	return s.league
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}
