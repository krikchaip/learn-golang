package testutils

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type testServer struct {
	*httptest.Server
}

func NewTestServer(h http.Handler) *testServer {
	// this starts up a HTTPS server which listens on a
	// randomly-chosen port of your local machine
	server := httptest.NewTLSServer(h)

	return &testServer{server}
}

func (ts *testServer) Get(
	t *testing.T,
	path string,
) (statusCode int, header http.Header, body string) {
	// server's dedicated client which configured to trust the server's TLS cert
	// and will be automatically closed upon server.Close()
	client := ts.Client()

	// we can get the server's listening URL by using server.URL,
	res, err := client.Get(fmt.Sprintf("%s%s", ts.URL, path))
	if err != nil {
		t.Fatal(err)
		return
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
		return
	}

	body = string(bytes.TrimSpace(b))
	defer res.Body.Close()

	statusCode = res.StatusCode
	header = res.Header

	return
}
