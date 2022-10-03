package server_test

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"

	server "22-building-application/controller/server"
	"22-building-application/entity"
	util "22-building-application/util/testing"

	"github.com/gorilla/websocket"
)

var (
	dummyGame = &util.GameSpy{}
)

func TestGETPlayers(t *testing.T) {
	store := util.NewStubPlayerStore(util.WithScores(map[string]int{
		"Pepper": 20,
		"Floyd":  10,
	}))
	sv := mustMakePlayerServer(t, store, dummyGame)

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
	sv := mustMakePlayerServer(t, store, dummyGame)

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
		sv := mustMakePlayerServer(t, store, dummyGame)

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
		sv := mustMakePlayerServer(t, store, dummyGame)

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
		sv := mustMakePlayerServer(t, store, dummyGame)

		req := util.NewGameRequest()
		res := httptest.NewRecorder()

		sv.ServeHTTP(res, req)

		util.AssertStatus(t, res.Code, http.StatusOK)
	})

	t.Run("start a game with 3 players, send some blind alerts down WS and declare Ruth the winner", func(t *testing.T) {
		winner := "Ruth"
		alertWith := "Blind is 100"

		store := util.NewStubPlayerStore()
		game := util.NewGameSpy(alertWith)

		server := httptest.NewServer(mustMakePlayerServer(t, store, game))
		defer server.Close()

		ws := mustDialWS(t, "ws"+strings.TrimPrefix(server.URL, "http")+"/ws")
		defer ws.Close()

		writeWSMessage(t, ws, "3")
		writeWSMessage(t, ws, winner)

		// ?? There is a delay between our WebSocket connection reading the message
		// ?? and recording the win and our test finishes before it happens.
		// ?? You can test this by putting a short time.Sleep before the final assertion.

		// TODO: refactor this
		time.Sleep(10 * time.Millisecond)

		util.AssertGameStartedWith(t, game, 3)
		util.AssertPlayerWin(t, store.GetWinCalls(), []string{winner})

		util.Within(t, 3*time.Second, func() {
			util.AssertWebsocketGotMsg(t, ws, alertWith)
		})
	})
}

func mustMakePlayerServer(
	t *testing.T,
	store entity.PlayerStore,
	game entity.Game,
) *server.PlayerServer {
	defer func() {
		if err := recover(); err != nil {
			t.Fatal("problem creating player server", err)
		}
	}()

	return server.NewPlayerServer(store, game)
}

func mustDialWS(t *testing.T, url string) *websocket.Conn {
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)

	if err != nil {
		t.Fatalf("could not open a ws connection on %s %v\n", url, err)
	}

	return ws
}

func writeWSMessage(t testing.TB, conn *websocket.Conn, msg string) {
	t.Helper()
	if err := conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
		t.Fatalf("could not send message over ws connection %v\n", err)
	}
}
