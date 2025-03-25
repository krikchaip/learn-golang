package main

import (
	"krikchaip/choose-your-own-adventure/pages"
	"net/http"
)

func main() {
	stories := loadJSON(args.Filepath)
	t := loadTemplate(pages.FS)

	http.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, args.Root, http.StatusSeeOther)
	})

	http.HandleFunc("/{arc}", func(w http.ResponseWriter, r *http.Request) {
		arc := r.PathValue("arc")
		s, ok := stories[arc]

		if !ok {
			http.NotFound(w, r)
			return
		}

		t.Execute(w, s)
	})

	serveHTTP(args.Addr)
}
