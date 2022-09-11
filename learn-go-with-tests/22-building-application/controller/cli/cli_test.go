package cli_test

import (
	"strings"
	"testing"

	"22-building-application/controller/cli"
	testutil "22-building-application/util/testing"
)

func TestCLI(t *testing.T) {
	t.Run("record Chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		store := testutil.NewStubPlayerStore()

		program := cli.NewPlayerCLI(store, in)
		program.PlayPoker()

		testutil.AssertPlayerWin(t, store.GetWinCalls(), []string{"Chris"})
	})

	t.Run("record Cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		store := testutil.NewStubPlayerStore()

		program := cli.NewPlayerCLI(store, in)
		program.PlayPoker()

		testutil.AssertPlayerWin(t, store.GetWinCalls(), []string{"Cleo"})
	})
}
