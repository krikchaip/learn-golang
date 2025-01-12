package main

import (
	"bytes"
	"io"
	"krikchaip/snippetbox/internal/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthz(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	healthz(rec, req)

	// you must call the Result() method
	// to get the actual response from the handler
	res := rec.Result()

	// check if the StatusCode is 200
	assert.Equal(t, res.StatusCode, http.StatusOK)

	// res.Body implements io.ReadCloser
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
		return
	}

	body = bytes.TrimSpace(body)
	assert.Equal(t, string(body), "OK")
}
