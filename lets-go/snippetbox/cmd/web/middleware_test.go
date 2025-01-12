package main

import (
	"io"
	"krikchaip/snippetbox/internal/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSecurityHeaders(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	// create a mock HTTP handler that we can pass to the 'securityHeaders' middleware
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	securityHeaders(next).ServeHTTP(rec, req)

	// you must call the Result() method
	// to get the actual response from the handler
	res := rec.Result()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	// check if the 'next' middleware is actually gets called
	assert.Equal(t, res.StatusCode, http.StatusOK)
	assert.Equal(t, strings.TrimSpace(string(body)), "OK")

	assert.Equal(
		t,
		res.Header.Get("Content-Security-Policy"),
		"default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com",
	)

	assert.Equal(t, res.Header.Get("Referrer-Policy"), "origin-when-cross-origin")
	assert.Equal(t, res.Header.Get("X-Content-Type-Options"), "nosniff")
	assert.Equal(t, res.Header.Get("X-Frame-Options"), "deny")
	assert.Equal(t, res.Header.Get("X-XSS-Protection"), "0")
}
