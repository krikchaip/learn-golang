package main

import (
	"fmt"
	"net/http"
)

func main() {
	stories := loadJSON(args.Filepath)

	http.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, args.Root, http.StatusSeeOther)
	})

	http.HandleFunc("/{arc}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, toJSON(stories[r.PathValue("arc")]))
	})

	serveHTTP(args.Addr)
}
