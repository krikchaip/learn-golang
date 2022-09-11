package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"21-acceptance-testing/lib/util"
	"22-building-application/controller/cli"
	"22-building-application/entity"
	"22-building-application/store"
)

const DB_SOURCE = "assets/game.db.json"

// go run 22-building-application/cmd/cli OR
// cd 22-building-application/cmd/cli && go run .
func main() {
	fmt.Println("Let's play poker!")
	fmt.Println(`Type "{Name} wins" to record a win`)

	st := setupStore()
	program := cli.NewPlayerCLI(st, os.Stdin)

	program.PlayPoker()
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
