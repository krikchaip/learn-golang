package store

import (
	"22-building-application/server"
	"encoding/json"
	"io"
)

// implements: server.PlayerStore
type FileSystemPlayerStore struct {
	source io.Reader
}

func NewFileSystemPlayerStore(source io.Reader) *FileSystemPlayerStore {
	return &FileSystemPlayerStore{source}
}

func (s *FileSystemPlayerStore) GetLeagueTable() (league []server.Player) {
	// ** reads and parse from source directly
	json.NewDecoder(s.source).Decode(&league)

	return league
}
