package store

import (
	"sync"

	"22-building-application/entity"
)

// implements: server.PlayerStore
type InMemoryPlayerStore struct {
	store map[string]int
	mut   sync.Mutex
}

func NewInMemoryPlayerStore() entity.PlayerStore {
	return &InMemoryPlayerStore{
		store: make(map[string]int),
	}
}

func (s *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return s.store[name]
}

func (s *InMemoryPlayerStore) RecordWin(name string) {
	s.mut.Lock()
	defer s.mut.Unlock()

	s.store[name]++
}

func (s *InMemoryPlayerStore) GetLeagueTable() (league entity.League) {
	for k, v := range s.store {
		league = append(league, entity.Player{Name: k, Wins: v})
	}

	return
}
