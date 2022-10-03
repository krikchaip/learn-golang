package testing

import (
	"io"
	"sync"
)

// implements: entity.Game
type GameSpy struct {
	startCalled bool
	startedWith int

	finishCalled bool
	finishedWith string

	mut sync.Mutex
}

// interface methods

func (g *GameSpy) Start(dest io.Writer, nPlayers int) {
	g.mut.Lock()
	defer g.mut.Unlock()

	g.startCalled = true
	g.startedWith = nPlayers
}

func (g *GameSpy) Finish(winner string) {
	g.mut.Lock()
	defer g.mut.Unlock()

	g.finishCalled = true
	g.finishedWith = winner
}

// helper methods

func (g *GameSpy) GetStartCalled() bool {
	g.mut.Lock()
	defer g.mut.Unlock()

	return g.startCalled
}

func (g *GameSpy) GetStartedWith() int {
	g.mut.Lock()
	defer g.mut.Unlock()

	return g.startedWith
}

func (g *GameSpy) GetFinishedWith() string {
	g.mut.Lock()
	defer g.mut.Unlock()

	return g.finishedWith
}
