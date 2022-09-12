package store

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"21-acceptance-testing/lib/util"
	"22-building-application/entity"
)

const (
	DB_SOURCE = "assets/game.db.json"
)

func SetupFileSystemStore() (store entity.PlayerStore, close func()) {
	// store := NewInMemoryPlayerStore()

	sourceName := resolvePath(DB_SOURCE)
	perm := os.O_CREATE | os.O_RDWR
	source, err := os.OpenFile(sourceName, perm, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", sourceName, err)
	}

	store = NewFileSystemPlayerStore(source)
	close = func() { source.Close() }

	return
}

func resolvePath(rel string) string {
	_, filename, _, _ := runtime.Caller(0)
	pkgRoot, err := util.FindRoot(filepath.Dir(filename))

	if err != nil {
		panic(err)
	}

	return filepath.Join(pkgRoot, rel)
}
