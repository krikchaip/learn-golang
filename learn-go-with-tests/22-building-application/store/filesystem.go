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
	for _, p := range s.GetLeagueTable() {
		if name == p.Name {
			return p.Wins
		}
	}

	return 0
}

func (s *FileSystemPlayerStore) GetLeagueTable() (league entity.League) {
	s.source.Seek(0, io.SeekStart)
	league, _ = entity.NewLeague(s.source)
	return
}

func (s *FileSystemPlayerStore) RecordWin(name string) {
	league := s.GetLeagueTable()

	for i, p := range league {
		if name == p.Name {
			// ** this will not work because when you `range` over a slice
			// ** you are returned a COPY OF AN ELEMENT at the current index.
			// p.Wins++

			// ** For that reason, we need to get the reference of the actual value
			// ** and then changing that value instead.
			league[i].Wins++
		}
	}

	// because the file cursor has already reached the end
	// from calling s.GetLeagueTable()
	s.source.Seek(0, io.SeekStart)

	json.NewEncoder(s.source).Encode(league)
}
