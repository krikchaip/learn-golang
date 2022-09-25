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

		testutil.AssertMessagesSentToUser(t, out, cli.PlayerPrompt)
		testutil.AssertGameStartedWith(t, game, 7)
	})

	t.Run("it prints an error when a non-numeric value is enterd and does not start the game", func(t *testing.T) {
		in := strings.NewReader("EIEI")
		out := &bytes.Buffer{}
		game := &testutil.GameSpy{}

		program := cli.NewPlayerCLI(in, out, game)
		program.PlayPoker()

		testutil.AssertGameNotStarted(t, game)
		testutil.AssertMessagesSentToUser(t, out, cli.PlayerPrompt+cli.BadPlayerInputErrMsg)
	})
}
