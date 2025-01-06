package main

import (
	"errors"
	"fmt"
	"krikchaip/snippetbox/internal/models"
	"net/http"
	"strconv"
)

func (app *application) defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from %q!", r.URL.Path)
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := app.newTemplateData(r)
	data.Snippets = snippets

	app.render(w, r, http.StatusOK, "home", data)
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	// retrieving path param
	id, err := strconv.Atoi(r.PathValue("id"))

	// field validation
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	snippet, err := app.snippets.Get(id)

	if errors.Is(err, models.ErrNoRecord) {
		http.NotFound(w, r)
		return
	}

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// fmt.Fprintf(w, "Display a specific snippet with ID %d", id)
	// fmt.Fprintf(w, "%+v", snippet)

	data := app.newTemplateData(r)
	data.Snippet = snippet

	app.render(w, r, http.StatusOK, "view", data)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	app.render(w, r, http.StatusOK, "create", data)
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	// parse form data received from the client.
	// the parsed data will be put into the r.PostForm and r.Form struct
	if err := r.ParseForm(); err != nil {
		app.cilentError(w, http.StatusBadRequest)
		return
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")

	// parse "expires" into an interger before using
	expires, err := strconv.Atoi(r.PostForm.Get("expires"))
	if err != nil {
		app.cilentError(w, http.StatusBadRequest)
		return
	}

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
