package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) cilentError(w http.ResponseWriter, statusCode int) {
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
