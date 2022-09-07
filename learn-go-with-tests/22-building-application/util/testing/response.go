package testing

import (
	"bytes"
	"encoding/json"
	tt "testing"

	"22-building-application/server"
)

func ParseLeagueFromResponse(
	t tt.TB,
	body *bytes.Buffer,
) (league []server.Player) {
	t.Helper()

	// ** transform JSON string into a struct from Buffer.Read()
	err := json.NewDecoder(body).Decode(&league)

	// // ?? alternative to json.Decoder (using json.Unmarshal)
	// err := json.Unmarshal(body.Bytes(), &league)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", body, err)
	}

	return
}
