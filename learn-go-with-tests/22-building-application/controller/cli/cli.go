package cli

import (
	"22-building-application/entity"
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	store   entity.PlayerStore
	scanner *bufio.Scanner
}

func NewPlayerCLI(store entity.PlayerStore, in io.Reader) *CLI {
	return &CLI{store, bufio.NewScanner(in)}
}

func (c *CLI) PlayPoker() {
	input := c.readLine()
	name, ok := extractWinner(input)

	if ok {
		c.store.RecordWin(name)
	}
}

func (c *CLI) readLine() string {
	c.scanner.Scan()
	return c.scanner.Text()
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
