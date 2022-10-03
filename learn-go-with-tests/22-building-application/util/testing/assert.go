package testing

import (
	"bytes"
	"net/http"
	"reflect"
	"strings"
	tt "testing"

	"22-building-application/entity"

	"github.com/gorilla/websocket"
)

func AssertStatus(t tt.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d want %d", got, want)
	}
}

func AssertContentJSON(t tt.TB, got http.Header) {
	t.Helper()
	if got.Get("content-type") != "application/json" {
		t.Errorf("response did not have content-type of application/json, got %v", got.Get("content-type"))
	}
}

func AssertResponseBody(t tt.TB, got *bytes.Buffer, want string) {
	t.Helper()
	if got.String() != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func AssertLeagueTable(t tt.TB, got, want []entity.Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func AssertScoreEquals(t tt.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func AssertNoPanic(t tt.TB, f func()) {
	t.Helper()

	defer func() {
		if err := recover(); err != nil {
			t.Fatalf("didn't expect a panic but got one, %v", err)
		}
	}()

	f()
}

func AssertPlayerWin(t tt.TB, got, want []string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("didn't record correct winner, got %v, want %v", got, want)
	}
}

func AssertScheduledAlert(t tt.TB, got, want ScheduleAlert) {
	t.Helper()
	if got.Amount != want.Amount {
		t.Errorf("got amount %d, want %d", got.Amount, want.Amount)
	}

	if got.At != want.At {
		t.Errorf("got scheduled time of %v, want %v", got.At, want.At)
	}
}

func AssertMessagesSentToUser(t tt.TB, out *bytes.Buffer, messages ...string) {
	t.Helper()

	got := out.String()
	want := strings.Join(messages, "")

	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}

func AssertGameStartedWith(t tt.TB, game *GameSpy, want int) {
	t.Helper()
	s := game.GetStartedWith()
	if s != want {
		t.Errorf("wanted Start called with %d but got %d", want, s)
	}
}

func AssertFinishCalledWith(t tt.TB, game *GameSpy, want string) {
	t.Helper()
	f := game.GetFinishedWith()
	if f != want {
		t.Errorf("wanted Finish called with %q but got %q", want, f)
	}
}

func AssertGameNotStarted(t tt.TB, game *GameSpy) {
	t.Helper()
	s := game.GetStartCalled()
	if s {
		t.Error("game should not have started")
	}
}

func AssertWebsocketGotMsg(t *tt.T, conn *websocket.Conn, want string) {
	t.Helper()

	_, msg, err := conn.ReadMessage()
	if err != nil {
		t.Fatalf("could not read message from ws connection %v\n", err)
	}

	if string(msg) != want {
		t.Errorf("got blind alert %q, want %q", msg, want)
	}
}
