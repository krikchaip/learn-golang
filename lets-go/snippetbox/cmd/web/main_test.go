// end-to-end testing
package main

import (
	"io"
	"krikchaip/snippetbox/internal/assert"
	"krikchaip/snippetbox/internal/models/mocks"
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

func TestSnippetViewE2E(t *testing.T) {
	app := newTestApplication(t)
	server := testutils.NewTestServer(t, app.routes())

	// must call Close() so that the server is shutdown when the test finishes
	defer server.Close()

	cases := []struct {
		name     string
		path     string
		wantCode int
		wantBody string
	}{
		{
			name:     "Valid ID",
			path:     "/snippet/view/1",
			wantCode: http.StatusOK,
			wantBody: "An old silent pond...",
		},
		{
			name:     "Non-existent ID",
			path:     "/snippet/view/2",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Negative ID",
			path:     "/snippet/view/-1",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Decimal ID",
			path:     "/snippet/view/1.23",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "String ID",
			path:     "/snippet/view/foo",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "Empty ID",
			path:     "/snippet/view/",
			wantCode: http.StatusNotFound,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			statusCode, _, body := server.Get(t, c.path)

			assert.Equal(t, statusCode, c.wantCode)

			if c.wantBody != "" {
				assert.StringContains(t, body, c.wantBody)
			}
		})
	}
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

		// mocked database layer
		snippets: mocks.NewSnippetModel(),
		users:    mocks.NewUserModel(),
	}

	return app
}
