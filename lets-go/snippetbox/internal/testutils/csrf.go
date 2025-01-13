package testutils

import (
	"html"
	"regexp"
	"testing"
)

// a pattern that captures the CSRF token value from the HTML for our user signup page
var csrfTokenRegex = regexp.MustCompile(`<input type='hidden' name='csrf_token' value='(.+)'>`)

func ExtractCSRFToken(t *testing.T, body string) string {
	// this returns the entire matched pattern in the first position,
	// and the values of any captured data in the subsequent positions
	matches := csrfTokenRegex.FindStringSubmatch(body)
	if len(matches) < 2 {
		t.Fatal("no csrf token found in body")
	}

	// the 'html/template' package automatically escapes all dynamically rendered data.
	// because the CSRF token is a base64 encoded string it will potentially
	// include the '+' character, and this will be escaped to &#43;
	return html.UnescapeString(matches[1])
}
