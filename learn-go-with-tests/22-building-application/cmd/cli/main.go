package main

import (
	"fmt"
	"os"

	"22-building-application/controller/cli"
	"22-building-application/store"
)

// go run 22-building-application/cmd/cli OR
// cd 22-building-application/cmd/cli && go run .
func main() {
	fmt.Println("Let's play poker!")
	fmt.Println(`Type "{Name} wins" to record a win`)

	st := store.SetupFileSystemStore()
	program := cli.NewPlayerCLI(st, os.Stdin)

	program.PlayPoker()
}
