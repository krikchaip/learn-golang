package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

// follows the pattern of "host:port"
// the absence of host means that the handler will listen to every host requested
// on the specific port
// const PORT = ":4000"

func main() {
	// ############ flags ###########################

	// define a command-line flag called "addr"
	addr := flag.String("addr", ":4000", "HTTP network address")

	// NOTE: Must be called after all flags are defined and before flags are accessed
	flag.Parse()

	// flag value is actally a pointer
	PORT := *addr

	// ############# dependencies ###################

	app := &application{
		// create a "structured logger" that writes to stdout in plain text
		logger: slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		})),
	}

	// ############# router #########################

	router := http.NewServeMux()

	// serve files out of the "./ui/static" directory
	fileServer := http.FileServer(http.Dir("ui/static"))

	// match a single slash, followed by nothing else (exact match)
	router.HandleFunc("GET /{$}", app.home)

	// this will match the specified pattern exactly
	router.HandleFunc("GET /snippet/view/{id}", app.snippetView)

	router.HandleFunc("GET /snippet/create", app.snippetCreate)
	router.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	// serves static files (subtree path pattern)
	// will match "/static/**", eg. "/static/css/main.css"
	router.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// a catch-all handler (subtree path pattern)
	// will match "/**", eg. "/foo", "/bar/bax/..."
	// router.HandleFunc("/", defaultHandler)

	// ############# start ##########################

	// log.Printf("starting server on %s", PORT)

	// logging with a structured logger (DEBUG > INFO > WARN > ERROR)
	// logger.Info("starting server", "PORT", PORT)
	app.logger.Info("starting server", slog.String("PORT", PORT))

	if err := http.ListenAndServe(PORT, router); err != nil {
		// log.Fatal(err)

		app.logger.Error(err.Error())
		os.Exit(1)
	}
}
