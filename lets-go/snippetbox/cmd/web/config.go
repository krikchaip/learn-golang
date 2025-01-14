package main

import "flag"

// follows the pattern of "host:port"
// the absence of host means that the handler will listen to every host requested
// on the specific port
// const PORT = ":4000"
var PORT string

// database source name (postgresql)
var DSN string

// tls certificate and its private key
var CERT_FILE, KEY_FILE string

// whether to enable debug mode
var DEBUG bool

func initializeFlags() {
	// define a command-line flag called "addr"
	// addr := flag.String("addr", ":4000", "HTTP network address")

	// flag value is actally a pointer
	// PORT = *addr

	// pointer-reference variation (directly assign value to PORT)
	flag.StringVar(&PORT, "addr", ":4000", "HTTP network address")

	flag.StringVar(
		&DSN,
		"dsn",
		"postgresql://web:secret@localhost:5432/snippetbox",
		"PostgreSQL data source name",
	)

	flag.StringVar(&CERT_FILE, "cert-file", "tls/cert.pem", "TLS certificate file")
	flag.StringVar(&KEY_FILE, "key-file", "tls/key.pem", "TLS certificate private key")

	flag.BoolVar(&DEBUG, "debug", false, "Enable debug mode")

	// NOTE: Must be called after all flags are defined and before flags are accessed
	flag.Parse()
}
