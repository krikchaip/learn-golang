package main

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gorilla/schema"
)

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()

		// capture a stack trace up until this point of calling
		trace = string(debug.Stack())
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri)

	var errMsg string

	// display full stack trace upon response when 'debug' mode is enabled
	if app.debug {
		errMsg = trace
	} else {
		errMsg = http.StatusText(http.StatusInternalServerError)
	}

	http.Error(w, errMsg, http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, statusCode int) {
	http.Error(w, http.StatusText(statusCode), statusCode)
}

func (app *application) render(
	w http.ResponseWriter,
	r *http.Request,
	status int,
	page string,
	data templateData,
) {
	// retrieve the appropriate template set from the cache
	// based on the page name (like 'home' page)
	t, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %q does not exist", page)
		app.serverError(w, r, err)
		return
	}

	buf := &bytes.Buffer{}

	// attempts to write the template to the buffer, instead of straight to the
	// w http.ResponseWriter. If there's an error, call our serverError() and return
	if err := t.ExecuteTemplate(buf, "base", data); err != nil {
		app.serverError(w, r, err)
		return
	}

	// Write out the provided HTTP status code ('200 OK', '400 Bad Request' etc).
	// NOTE: MUST BE executed before any call to Write()
	w.WriteHeader(status)

	// If the template is written to the buffer without any errors
	// we are safe to go ahead and write the content to the http.ResponseWriter
	buf.WriteTo(w)
}

// 'dst' must be a pointer to a struct
func (app *application) decodePostForm(r *http.Request, dst any) error {
	// parse form data received from the client.
	// the parsed data will be put into the r.PostForm and r.Form struct
	if err := r.ParseForm(); err != nil {
		return err
	}

	// parse form data using a 3rd party package
	if err := app.decoder.Decode(dst, r.PostForm); err != nil {
		// NOTE: errors.As() requires the second parameter to be a pointer
		var conversionError *schema.ConversionError

		// check for this specific error type and panic()
		// instead of returning the error to the user
		if errors.As(err, &conversionError) {
			panic(err)
		}

		// 'nosurf' middleware requires HTML form to store
		// an additional 'csrf_token' key for CSRF protection.
		// we'll skip checking the field
		if err, ok := err.(schema.MultiError); ok && err["csrf_token"] != nil {
			return nil
		}

		return err
	}

	return nil
}

func (app *application) isAuthenticated(r *http.Request) bool {
	isAuthenticated, ok := r.Context().Value(isAuthenticatedContextKey).(bool)

	if !ok {
		return false
	}

	return isAuthenticated
}
