package cli

import (
	"22-building-application/entity"
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	store entity.PlayerStore
	in    io.Reader
}

func NewPlayerCLI(store entity.PlayerStore, in io.Reader) *CLI {
	return &CLI{store, in}
}

func (c *CLI) PlayPoker() {
	scanner := bufio.NewScanner(c.in)

	scanner.Scan()
	name, ok := extractWinner(scanner.Text())

	if ok {
		c.store.RecordWin(name)
	}
}

func extractWinner(input string) (winner string, ok bool) {
	xs := strings.SplitN(input, " ", 2)
	name, cmd := xs[0], xs[1]

	if cmd != "wins" {
		ok = false
		return
	}

	winner = name
	ok = true

	return
}
