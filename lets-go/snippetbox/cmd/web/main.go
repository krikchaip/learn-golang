package main

import (
	"database/sql"
	"krikchaip/snippetbox/internal/models"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	// create a "structured logger" that writes to stdout in plain text
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))

	// initialize database connection pool
	db, err := openDB(DSN)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	// initialize html template cache
	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// application instance with all dependencies setup
	app := &application{
		logger:        logger,
		snippets:      models.NewSnippetModel(db),
		templateCache: templateCache,
	}

	// logging with the default logger
	// log.Printf("starting server on %s", PORT)

	// logging with a structured logger (DEBUG > INFO > WARN > ERROR)
	// logger.Info("starting server", "PORT", PORT)
	logger.Info("starting server", slog.String("PORT", PORT))

	if err := http.ListenAndServe(PORT, app.routes()); err != nil {
		// log.Fatal(err)

		// 'slog' does not have .Fatal(), so we have to do it ourselves
		logger.Error(err.Error())
		os.Exit(1)
	}
}

// wraps sql.Open() function and reutrns a connection pool
func openDB(dsn string) (*sql.DB, error) {
	// initialize a connection pool for future use
	// (no connections are actually created at this step)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	// connects to the db and check for errors
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
