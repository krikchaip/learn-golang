package testing

import (
	"bytes"
	tt "testing"

	"22-building-application/server"
)

func ParseLeagueFromResponse(
	t tt.TB,
	body *bytes.Buffer,
) (league []server.Player) {
	t.Helper()

	league, err := server.NewLeague(body)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", body, err)
	}

	return
}
