package main

import (
	"log/slog"
	"net/http"
)

func (app *application) newHTTPServer() *http.Server {
	server := &http.Server{
		// pass 'Addr' and 'Handler' as with http.ListenAndServe
		Addr:    PORT,
		Handler: app.routes(),

		// converts our structured logger (slog) into a *log.Logger
		// which write log entries at 'Error' level
		ErrorLog: slog.NewLogLogger(app.logger.Handler(), slog.LevelError),
	}

	return server
}
