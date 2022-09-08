package store

import (
	"22-building-application/entity"
	"encoding/json"
	"io"
)

// implements: server.PlayerStore
type FileSystemPlayerStore struct {
	source io.ReadWriteSeeker
}

func NewFileSystemPlayerStore(source io.ReadWriteSeeker) entity.PlayerStore {
	return &FileSystemPlayerStore{source}
}

func (s *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := s.GetLeagueTable().Find(name)

	if player == nil {
		return 0
	}

	return player.Wins
}

func (s *FileSystemPlayerStore) GetLeagueTable() (league entity.League) {
	s.source.Seek(0, io.SeekStart)
	league, _ = entity.NewLeague(s.source)
	return
}

func (s *FileSystemPlayerStore) RecordWin(name string) {
	league := s.GetLeagueTable()
	player := league.Find(name)

	if player != nil {
		player.Wins++
	}

	// because the file cursor has already reached the end
	// from calling s.GetLeagueTable()
	s.source.Seek(0, io.SeekStart)

	json.NewEncoder(s.source).Encode(league)
}
