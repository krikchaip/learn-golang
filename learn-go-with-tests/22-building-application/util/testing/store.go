package testing

import "22-building-application/entity"

// implements: entity.PlayerStore
type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   entity.League
}

type StubPlayerStoreOption = func(*StubPlayerStore)

func NewStubPlayerStore(options ...StubPlayerStoreOption) *StubPlayerStore {
	store := &StubPlayerStore{
		scores:   map[string]int{},
		winCalls: []string{},
		league:   entity.League{},
	}

	for _, opt := range options {
		opt(store)
	}

	return store
}

// interface methods

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) GetLeagueTable() entity.League {
	return s.league
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

// helper methods

func (s *StubPlayerStore) GetWinCalls() []string {
	return s.winCalls
}

func WithScores(scores map[string]int) StubPlayerStoreOption {
	return func(store *StubPlayerStore) {
		store.scores = scores
	}
}

func WithLeague(league entity.League) StubPlayerStoreOption {
	return func(store *StubPlayerStore) {
		store.league = league
	}
}
