package store

import (
	"22-building-application/entity"
	"encoding/json"
	"io"
	"sync"
)

// implements: server.PlayerStore
type FileSystemPlayerStore struct {
	source io.ReadWriteSeeker
	cache  entity.League
	mut    sync.Mutex
}

func NewFileSystemPlayerStore(source io.ReadWriteSeeker) entity.PlayerStore {
	source.Seek(0, io.SeekStart)

	// initialize cache to improve performance
	cache, _ := entity.NewLeague(source)

	return &FileSystemPlayerStore{
		source: source,
		cache:  cache,
	}
}

func (s *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := s.cache.Find(name)

	if player == nil {
		return 0
	}

	return player.Wins
}

func (s *FileSystemPlayerStore) GetLeagueTable() (league entity.League) {
	league = s.cache
	return
}

func (s *FileSystemPlayerStore) RecordWin(name string) {
	s.mut.Lock()
	defer s.mut.Unlock()

	player := s.cache.Find(name)

	if player != nil {
		player.Wins++
	} else {
		s.cache = append(s.cache, entity.Player{Name: name, Wins: 1})
	}

	// because the file cursor has already reached the end
	// from calling s.GetLeagueTable()
	s.source.Seek(0, io.SeekStart)

	json.NewEncoder(s.source).Encode(s.cache)
}
