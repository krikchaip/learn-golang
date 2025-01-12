// end-to-end testing
package main

import (
	"io"
	"krikchaip/snippetbox/internal/assert"
	"krikchaip/snippetbox/internal/testutils"
	"log/slog"
	"net/http"
	"testing"
)

func TestHealthzE2E(t *testing.T) {
	app := newTestApplication()
	server := testutils.NewTestServer(app.routes())

	// must call Close() so that the server is shutdown when the test finishes
	defer server.Close()

	statusCode, _, body := server.Get(t, "/healthz")

	assert.Equal(t, statusCode, http.StatusOK)
	assert.Equal(t, body, "OK")
}

func newTestApplication() *application {
	// mock dependencies
	app := &application{
		// will discard anything written to io.Discard
		logger: slog.New(slog.NewTextHandler(io.Discard, nil)),
	}

	return app
}
