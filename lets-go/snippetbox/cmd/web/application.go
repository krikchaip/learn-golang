package main

import "log/slog"

// the application's dependencies (for dependency injection)
type application struct {
	logger *slog.Logger
}
