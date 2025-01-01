package main

import (
	"database/sql"
	"log/slog"
)

// the application's dependencies (for dependency injection)
type application struct {
	logger *slog.Logger
	db     *sql.DB
}
