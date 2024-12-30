package main

import (
	"log/slog"
	"net/http"
	"os"
)

func main() {
	app := &application{
		// create a "structured logger" that writes to stdout in plain text
		logger: slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		})),
	}

	// logging with the default logger
	// log.Printf("starting server on %s", PORT)

	// logging with a structured logger (DEBUG > INFO > WARN > ERROR)
	// logger.Info("starting server", "PORT", PORT)
	app.logger.Info("starting server", slog.String("PORT", PORT))

	if err := http.ListenAndServe(PORT, app.routes()); err != nil {
		// log.Fatal(err)

		// 'slog' does not have .Fatal(), so we have to do it ourselves
		app.logger.Error(err.Error())
		os.Exit(1)
	}
}
