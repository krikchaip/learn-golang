package main

import (
	"crypto/tls"
	"log/slog"
	"net/http"
	"time"
)

func (app *application) newHTTPServer() *http.Server {
	tlsConfig := &tls.Config{
		// changing the curve preferences value, so that only elliptic curves with
		// **assembly** implementations are used
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	server := &http.Server{
		// pass 'Addr' and 'Handler' as with http.ListenAndServe
		Addr:    PORT,
		Handler: app.routes(),

		// converts our structured logger (slog) into a *log.Logger
		// which write log entries at 'Error' level
		ErrorLog: slog.NewLogLogger(app.logger.Handler(), slog.LevelError),

		// customized tls config for running HTTPs
		TLSConfig: tlsConfig,

		// sets the 'keep-alives' timeout, helps reduce latency because a client
		// can reuse the same connection for multiple requests
		// without having to repeat the TLS handshake
		IdleTimeout: 1 * time.Minute,

		// closes the connection if the request headers or body are still being read
		// 'n' seconds after the request is first accepted, helps prevent the 'Slowloris' attacks
		ReadTimeout: 5 * time.Second,

		// will close the underlying connection if our server
		// attempts to write to the connection after a given period
		//   - for HTTP, starts countdown after w.Write() and finished reading the request header
		//   - for HTTPs, starts countdown after w.Write() and the request is first accepted
		WriteTimeout: 10 * time.Second,

		// limit the maximum size of the request header, default to '1MB'
		// if the header size is exceeded, returns '403 Request Header Fields Too Large'
		// ref: https://github.com/golang/go/blob/4b36e129f865f802eb87f7aa2b25e3297c5d8cfd/src/net/http/server.go#L871
		MaxHeaderBytes: 524288,
	}

	return server
}
