package cli_test

import (
	"bytes"
	"strings"
	"testing"

	"22-building-application/controller/cli"
	testutil "22-building-application/util/testing"
)

func TestCLI(t *testing.T) {
	t.Run("it prompts the user to enter the number of players and starts the game", func(t *testing.T) {
		in := strings.NewReader("7\n")
		out := &bytes.Buffer{}
		game := &testutil.GameSpy{}

		program := cli.NewPlayerCLI(in, out, game)
		program.PlayPoker()

		gotPrompt := out.String()
		wantPrompt := cli.PlayerPrompt

		if gotPrompt != wantPrompt {
			t.Errorf("got %q want %q", gotPrompt, wantPrompt)
		}

		if game.StartedWith != 7 {
			t.Errorf("wanted Start called with 7 but got %d", game.StartedWith)
		}
	})

	t.Run("it prints an error when a non-numeric value is enterd and does not start the game", func(t *testing.T) {
		in := strings.NewReader("EIEI")
		out := &bytes.Buffer{}
		game := &testutil.GameSpy{}

		program := cli.NewPlayerCLI(in, out, game)
		program.PlayPoker()

		if game.StartCalled {
			t.Error("game should not have started")
		}
	})
}
