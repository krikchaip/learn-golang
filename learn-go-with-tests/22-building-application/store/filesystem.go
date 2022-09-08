package store

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"sync"

	"22-building-application/entity"
	"22-building-application/util"
)

// implements: server.PlayerStore
type FileSystemPlayerStore struct {
	db    *json.Encoder
	cache entity.League
	mut   sync.Mutex
}

func NewFileSystemPlayerStore(source *os.File) entity.PlayerStore {
	initializeJSONFile(source)

	tape := util.NewTape(source)

	// initialize cache to improve performance
	cache, err := entity.NewLeague(tape)

	if err != nil {
		panic(fmt.Errorf("problem loading player store from file %s, %v", source.Name(), err))
	}

	return &FileSystemPlayerStore{
		db:    json.NewEncoder(tape),
		cache: cache,
	}
}

func (s *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := s.cache.Find(name)

	if player == nil {
		return 0
	}

	return player.Wins
}

func (s *FileSystemPlayerStore) GetLeagueTable() entity.League {
	// ** sorting a slice using "sort" package
	sort.Slice(s.cache, func(i, j int) bool {
		return s.cache[i].Wins > s.cache[j].Wins
	})

	return s.cache
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

	s.db.Encode(s.cache)
}

func initializeJSONFile(file *os.File) {
	file.Seek(0, io.SeekStart)
	defer file.Seek(0, io.SeekStart)

	info, err := file.Stat()

	if err != nil {
		panic(fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err))
	}

	if info.Size() == 0 {
		file.WriteString("[]")
	}
}
