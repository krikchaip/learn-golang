package testutils

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"net/url"
	"testing"
)

type testServer struct {
	*httptest.Server
}

func NewTestServer(t *testing.T, h http.Handler) *testServer {
	// this starts up a HTTPS server which listens on a
	// randomly-chosen port of your local machine
	server := httptest.NewTLSServer(h)

	jar, err := cookiejar.New(nil)
	if err != nil {
		t.Fatal(err)
	}

	// add the cookie jar to the test server client. Any response cookies will
	// now be stored and sent with subsequent requests when using this client
	server.Client().Jar = jar

	// disable redirect-following for the test server client,
	// returning a http.ErrUseLastResponse error forces the client
	// to immediately return the received response
	server.Client().CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

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

func (ts *testServer) PostForm(
	t *testing.T,
	path string,
	data url.Values,
) (statusCode int, header http.Header, body string) {
	client := ts.Client()

	res, err := client.PostForm(fmt.Sprintf("%s%s", ts.URL, path), data)
	if err != nil {
		t.Fatal(err)
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
