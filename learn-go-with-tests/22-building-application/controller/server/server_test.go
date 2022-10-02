package server_test

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	server "22-building-application/controller/server"
	"22-building-application/entity"
	util "22-building-application/util/testing"

	"github.com/gorilla/websocket"
)

func TestGETPlayers(t *testing.T) {
	store := util.NewStubPlayerStore(util.WithScores(map[string]int{
		"Pepper": 20,
		"Floyd":  10,
	}))
	sv := server.NewPlayerServer(store)

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
	store := util.NewStubPlayerStore()
	sv := server.NewPlayerServer(store)

	t.Run("it returns accepted on POST", func(t *testing.T) {
		req := util.NewPostWinRequest("Pepper")
		res := httptest.NewRecorder()

		sv.ServeHTTP(res, req)

		util.AssertStatus(t, res.Code, http.StatusAccepted)

		got := store.GetWinCalls()
		want := []string{"Pepper"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v calls to RecordWin want %v", got, want)
		}
	})
}

func TestLeague(t *testing.T) {
	t.Run("it returns 200 on /league", func(t *testing.T) {
		store := util.NewStubPlayerStore()
		sv := server.NewPlayerServer(store)

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

		store := util.NewStubPlayerStore(util.WithLeague(wantedLeague))
		sv := server.NewPlayerServer(store)

		req := util.NewLeagueRequest()
		res := httptest.NewRecorder()

		sv.ServeHTTP(res, req)

		got := util.ParseLeagueFromResponse(t, res.Body)
		util.AssertStatus(t, res.Code, http.StatusOK)
		util.AssertLeagueTable(t, got, wantedLeague)
		util.AssertContentJSON(t, res.Result().Header)
	})
}

func TestGame(t *testing.T) {
	t.Run("GET /game returns 200", func(t *testing.T) {
		store := util.NewStubPlayerStore()
		sv := server.NewPlayerServer(store)

		req := util.NewGameRequest()
		res := httptest.NewRecorder()

		sv.ServeHTTP(res, req)

		util.AssertStatus(t, res.Code, http.StatusOK)
	})

	t.Run("when we get a message over a websocket. it is a winner of a game", func(t *testing.T) {
		winner := "Ruth"

		store := util.NewStubPlayerStore()
		sv := server.NewPlayerServer(store)

		server := httptest.NewServer(sv)
		defer server.Close()

		url := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"

		ws, _, err := websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			t.Fatalf("could not open a ws connection on %s %v", url, err)
		}
		defer ws.Close()

		if err := ws.WriteMessage(websocket.TextMessage, []byte(winner)); err != nil {
			t.Fatalf("could not send message over ws connection %v", err)
		}

		util.AssertPlayerWin(t, store.GetWinCalls(), []string{winner})
	})
}
