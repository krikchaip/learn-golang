package testing

import (
	"bytes"
	tt "testing"

	"22-building-application/entity"
)

func ParseLeagueFromResponse(
	t tt.TB,
	body *bytes.Buffer,
) (league []entity.Player) {
	t.Helper()

	league, err := entity.NewLeague(body)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", body, err)
	}

	return
}
