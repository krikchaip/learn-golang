package main

import (
	"errors"
	"fmt"
	"krikchaip/snippetbox/internal/models"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"unicode/utf8"
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
		app.clientError(w, http.StatusBadRequest)
		return
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")

	// parse "expires" into an interger before using
	expires, err := strconv.Atoi(r.PostForm.Get("expires"))
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// validation cases:
	//   - Check that the "title" and "content" fields are not empty
	//   - Check that the "title" field is not more than 100 characters long
	//   - Check that the "expires" value exactly matches one of our permitted values (1, 7 or 365 days)
	fieldErrors := make(map[string]string)

	if strings.TrimSpace(title) == "" {
		fieldErrors["title"] = "This field cannot be blank"
	} else if utf8.RuneCountInString(title) > 100 {
		fieldErrors["title"] = "This field cannot be more than 100 characters long"
	}

	if strings.TrimSpace(content) == "" {
		fieldErrors["content"] = "This field cannot be blank"
	}

	if !slices.Contains([]int{1, 7, 365}, expires) {
		fieldErrors["expires"] = "This field must equal 1, 7 or 365"
	}

	// if there are any errors, dump them in a plain text HTTP response
	if len(fieldErrors) > 0 {
		fmt.Fprint(w, fieldErrors)
		return
	}

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
