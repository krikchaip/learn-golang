package entity

import (
	"io"
	"time"
)

type Game interface {
	Start(dest io.Writer, nPlayers int)
	Finish(winner string)
}

// implements: entity.Game
type TexasHoldem struct {
	store   PlayerStore
	alerter BlindAlerter
}

func NewTexasHoldem(alerter BlindAlerter, store PlayerStore) *TexasHoldem {
	return &TexasHoldem{store, alerter}
}

func (p *TexasHoldem) Start(dest io.Writer, nPlayers int) {
	blindIncrement := time.Duration(5+nPlayers) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second

	for _, b := range blinds {
		p.alerter.ScheduleAlertAt(dest, blindTime, b)
		blindTime = blindTime + blindIncrement
	}
}

func (p *TexasHoldem) Finish(winner string) {
	p.store.RecordWin(winner)
}
