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
	// I was trying to define the CLI flags in the init() function,
	// but it seemed not getting along with the 'go test' command
	// ref: https://stackoverflow.com/questions/64704124/flag-provided-but-not-defined-error-in-go-test-despite-the-flag-being-defined
	//      https://stackoverflow.com/questions/29699982/go-test-flag-flag-provided-but-not-defined
	initializeFlags()

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

	// cookies will only be sent when HTTPs is used
	sessionManager.Cookie.Secure = true

	// set this option to 'Strict' if your application will potentially
	// have other websites linking to it. Otherwise, set it to 'Lax'
	sessionManager.Cookie.SameSite = http.SameSiteStrictMode

	// application instance with all dependencies setup
	app := &application{
		debug: DEBUG,

		logger:         logger,
		templateCache:  templateCache,
		decoder:        decoder,
		sessionManager: sessionManager,

		snippets: models.NewSnippetModel(db),
		users:    models.NewUserModel(db),
	}

	// initialize a new http.Server struct to replace http.ListenAndServe
	server := app.newHTTPServer()

	// logging with the default logger
	// log.Printf("starting server on %s", PORT)

	// logging with a structured logger (DEBUG > INFO > WARN > ERROR)
	// logger.Info("starting server", "PORT", PORT)
	logger.Info("starting server", slog.String("PORT", PORT))

	// be sure to execute 'make gen-cert' before starting the server!
	if err := server.ListenAndServeTLS(CERT_FILE, KEY_FILE); err != nil {
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
