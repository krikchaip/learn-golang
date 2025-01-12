// end-to-end testing
package main

import (
	"bytes"
	"fmt"
	"io"
	"krikchaip/snippetbox/internal/assert"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthzE2E(t *testing.T) {
	// mock dependencies
	app := &application{
		// will discard anything written to io.Discard
		logger: slog.New(slog.NewTextHandler(io.Discard, nil)),
	}

	// this starts up a HTTPS server which listens on a
	// randomly-chosen port of your local machine
	server := httptest.NewTLSServer(app.routes())

	// must call Close() so that the server is shutdown when the test finishes
	defer server.Close()

	// server's dedicated client which configured to trust the server's TLS cert
	// and will be automatically closed upon server.Close()
	client := server.Client()

	// we can get the server's listening URL by using server.URL,
	res, err := client.Get(fmt.Sprintf("%s/healthz", server.URL))
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(t, res.StatusCode, http.StatusOK)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
		return
	}

	body = bytes.TrimSpace(body)
	defer res.Body.Close()

	assert.Equal(t, string(body), "OK")
}
