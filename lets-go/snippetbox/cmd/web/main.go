package main

import (
	"context"
	"krikchaip/snippetbox/internal/models"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/pgxstore"
	"github.com/alexedwards/scs/v2"
	"github.com/gorilla/schema"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

// it is recommended to set a decoder instance as a package global
// ref: https://github.com/gorilla/schema
var decoder = schema.NewDecoder()

func main() {
	// create a "structured logger" that writes to stdout in plain text
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))

	// initialize database connection pool
	pool, err := openDBPool(DSN)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// converts pgx database pool into a standard sql.DB interface
	db := stdlib.OpenDBFromPool(pool)

	defer pool.Close()
	defer db.Close()

	// initialize html template cache
	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// using a 3rd party user session manager
	// ref: https://github.com/alexedwards/scs
	sessionManager := scs.New()
	sessionManager.Store = pgxstore.New(pool)
	sessionManager.Lifetime = 12 * time.Hour

	// application instance with all dependencies setup
	app := &application{
		logger:         logger,
		snippets:       models.NewSnippetModel(db),
		templateCache:  templateCache,
		decoder:        decoder,
		sessionManager: sessionManager,
	}

	// initialize a new http.Server struct to replace http.ListenAndServe
	server := &http.Server{
		// pass 'Addr' and 'Handler' as with http.ListenAndServe
		Addr:    PORT,
		Handler: app.routes(),

		// converts our structured logger (slog) into a *log.Logger
		// which write log entries at 'Error' level
		ErrorLog: slog.NewLogLogger(app.logger.Handler(), slog.LevelError),
	}

	// logging with the default logger
	// log.Printf("starting server on %s", PORT)

	// logging with a structured logger (DEBUG > INFO > WARN > ERROR)
	// logger.Info("starting server", "PORT", PORT)
	logger.Info("starting server", slog.String("PORT", PORT))

	if err := server.ListenAndServe(); err != nil {
		// log.Fatal(err)

		// 'slog' does not have .Fatal(), so we have to do it ourselves
		logger.Error(err.Error())
		os.Exit(1)
	}
}

// wraps db init functions and returns a connection pool
func openDBPool(dsn string) (*pgxpool.Pool, error) {
	ctx := context.Background()

	// initialize a connection pool for future use
	// (no connections are actually created at this step)
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	// connects to the db and check for errors
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}

	return pool, nil
}
