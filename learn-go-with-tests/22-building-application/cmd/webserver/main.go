package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"21-acceptance-testing/lib/util"
	"22-building-application/controller/server"
	"22-building-application/entity"
	"22-building-application/store"
)

const DB_SOURCE = "assets/game.db.json"

// go run 22-building-application OR
// cd 22-building-application && go run .
func main() {
	st := setupStore()
	sv := server.NewPlayerServer(st)

	// we wrap the call in log.Fatal
	// just in case if there is a problem with ListenAndServe.
	// eg. port already being used, etc.
	log.Fatal(http.ListenAndServe(":3000", sv))
}

func setupStore() entity.PlayerStore {
	// store := store.NewInMemoryPlayerStore()

	sourceName := resolvePath(DB_SOURCE)
	perm := os.O_CREATE | os.O_RDWR
	source, err := os.OpenFile(sourceName, perm, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", sourceName, err)
	}

	store := store.NewFileSystemPlayerStore(source)

	return store
}

func resolvePath(rel string) string {
	_, filename, _, _ := runtime.Caller(0)
	pkgRoot, err := util.FindRoot(filepath.Dir(filename))

	if err != nil {
		panic(err)
	}

	return filepath.Join(pkgRoot, rel)
}
