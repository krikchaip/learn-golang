// end-to-end testing
package main

import (
	"io"
	"krikchaip/snippetbox/internal/assert"
	"krikchaip/snippetbox/internal/models/mocks"
	"krikchaip/snippetbox/internal/testutils"
	"log/slog"
	"net/http"
	"net/url"
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

func TestSnippetCreateE2E(t *testing.T) {
	app := newTestApplication(t)
	server := testutils.NewTestServer(t, app.routes())

	// must call Close() so that the server is shutdown when the test finishes
	defer server.Close()

	t.Run("Unauthenticated", func(t *testing.T) {
		code, header, _ := server.Get(t, "/snippet/create")

		assert.Equal(t, code, http.StatusSeeOther)
		assert.Equal(t, header.Get("Location"), "/user/login")
	})

	t.Run("Authenticated", func(t *testing.T) {
		_, _, body := server.Get(t, "/user/login")

		csrfToken := testutils.ExtractCSRFToken(t, body)
		t.Logf("CSRF token is: %q", csrfToken)

		data := url.Values{}
		data.Add("email", mocks.MockUser.Email)
		data.Add("password", mocks.MockUser.Password)
		data.Add("csrf_token", csrfToken)

		server.PostForm(t, "/user/login", data)
		code, _, body := server.Get(t, "/snippet/create")

		assert.Equal(t, code, http.StatusOK)
		assert.StringContains(t, body, `<form action="/snippet/create" method="POST">`)
	})
}

func TestUserSignupE2E(t *testing.T) {
	app := newTestApplication(t)
	server := testutils.NewTestServer(t, app.routes())

	// must call Close() so that the server is shutdown when the test finishes
	defer server.Close()

	// make the GET request and then extract the CSRF token from the response body
	_, _, body := server.Get(t, "/user/signup")
	csrfToken := testutils.ExtractCSRFToken(t, body)

	// log the token specifically in our test output, an alternative to fmt.Printf()
	t.Logf("CSRF token is: %q", csrfToken)

	formTag := `<form action="/user/signup" method="POST" novalidate>`

	cases := []struct {
		name         string
		userName     string
		userEmail    string
		userPassword string
		csrfToken    string
		wantCode     int
		wantFormTag  string
	}{
		{
			name:         "Valid submission",
			userName:     mocks.MockUser.Name,
			userEmail:    mocks.MockUser.Email,
			userPassword: mocks.MockUser.Password,
			csrfToken:    csrfToken,
			wantCode:     http.StatusSeeOther,
		},
		{
			name:         "Invalid CSRF Token",
			userName:     mocks.MockUser.Name,
			userEmail:    mocks.MockUser.Email,
			userPassword: mocks.MockUser.Password,
			csrfToken:    "wrongToken",
			wantCode:     http.StatusBadRequest,
		},
		{
			name:         "Empty name",
			userName:     "",
			userEmail:    mocks.MockUser.Email,
			userPassword: mocks.MockUser.Password,
			csrfToken:    csrfToken,
			wantCode:     http.StatusUnprocessableEntity,
			wantFormTag:  formTag,
		},
		{
			name:         "Empty email",
			userName:     mocks.MockUser.Name,
			userEmail:    "",
			userPassword: mocks.MockUser.Password,
			csrfToken:    csrfToken,
			wantCode:     http.StatusUnprocessableEntity,
			wantFormTag:  formTag,
		},
		{
			name:         "Empty password",
			userName:     mocks.MockUser.Name,
			userEmail:    mocks.MockUser.Email,
			userPassword: "",
			csrfToken:    csrfToken,
			wantCode:     http.StatusUnprocessableEntity,
			wantFormTag:  formTag,
		},
		{
			name:         "Invalid email",
			userName:     mocks.MockUser.Name,
			userEmail:    "bob@example.",
			userPassword: mocks.MockUser.Password,
			csrfToken:    csrfToken,
			wantCode:     http.StatusUnprocessableEntity,
			wantFormTag:  formTag,
		},
		{
			name:         "Short password",
			userName:     mocks.MockUser.Name,
			userEmail:    mocks.MockUser.Email,
			userPassword: "pa$$",
			csrfToken:    csrfToken,
			wantCode:     http.StatusUnprocessableEntity,
			wantFormTag:  formTag,
		},
		{
			name:         "Duplicate email",
			userName:     mocks.MockUser.Name,
			userEmail:    mocks.DupeEmail,
			userPassword: mocks.MockUser.Password,
			csrfToken:    csrfToken,
			wantCode:     http.StatusUnprocessableEntity,
			wantFormTag:  formTag,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			data := url.Values{}

			data.Add("name", c.userName)
			data.Add("email", c.userEmail)
			data.Add("password", c.userPassword)
			data.Add("csrf_token", c.csrfToken)

			statusCode, _, body := server.PostForm(t, "/user/signup", data)

			assert.Equal(t, statusCode, c.wantCode)

			if c.wantFormTag != "" {
				assert.StringContains(t, body, c.wantFormTag)
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
