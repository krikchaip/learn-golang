package entity

import "time"

type Game struct {
	store   PlayerStore
	alerter BlindAlerter
}

func NewGame(alerter BlindAlerter, store PlayerStore) *Game {
	return &Game{store, alerter}
}

func (p *Game) Start(nPlayers int) {
	blindIncrement := time.Duration(5+nPlayers) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second

	for _, b := range blinds {
		p.alerter.ScheduleAlertAt(blindTime, b)
		blindTime = blindTime + blindIncrement
	}
}

func (p *Game) Finish(winner string) {
	p.store.RecordWin(winner)
}
