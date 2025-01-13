// end-to-end testing
package main

import (
	"io"
	"krikchaip/snippetbox/internal/assert"
	"krikchaip/snippetbox/internal/testutils"
	"log/slog"
	"net/http"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/gorilla/schema"
)

func TestHealthzE2E(t *testing.T) {
	app := newTestApplication(t)
	server := testutils.NewTestServer(t, app.routes())

	// must call Close() so that the server is shutdown when the test finishes
	defer server.Close()

	statusCode, _, body := server.Get(t, "/healthz")

	assert.Equal(t, statusCode, http.StatusOK)
	assert.Equal(t, body, "OK")
}

func newTestApplication(t *testing.T) *application {
	// will discard anything written to io.Discard
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))

	templateCache, err := newTemplateCache()
	if err != nil {
		t.Fatal(err)
	}

	decoder = schema.NewDecoder()

	// if no store is set, the SCS package will default to using a transient in-memory store
	sessionManager := scs.New()
	sessionManager.Store = nil
	sessionManager.Lifetime = 12 * time.Hour
	sessionManager.Cookie.Secure = true

	// mock dependencies
	app := &application{
		logger:         logger,
		templateCache:  templateCache,
		decoder:        decoder,
		sessionManager: sessionManager,
	}

	return app
}