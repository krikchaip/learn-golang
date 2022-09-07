package testing

import (
	"bytes"
	"net/http"
	"reflect"
	tt "testing"

	"22-building-application/server"
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

func AssertLeagueTable(t tt.TB, got, want []server.Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}