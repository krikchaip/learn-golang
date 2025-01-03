package main

import (
	"errors"
	"fmt"
	"krikchaip/snippetbox/internal/models"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// add a new header key
	// NOTE: MUST BE executed before any call to WriteHeader() or Write()
	w.Header().Add("X-App-Env", "development")

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

func (app *application) snippetCreate(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)

	// send a 201 status code.
	// NOTE: MUST BE executed before any call to Write()
	// w.WriteHeader(http.StatusCreated)

	// w.Write([]byte("Save a new snippet..."))
}

func (app *application) defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from %q!", r.URL.Path)
}
