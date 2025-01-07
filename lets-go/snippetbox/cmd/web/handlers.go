package main

import (
	"errors"
	"fmt"
	"krikchaip/snippetbox/internal/models"
	"krikchaip/snippetbox/internal/validator"
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

	// initialize template data so that it won't complain of missing values
	data.Form = snippetCreateForm{
		Expires: 365,
	}

	app.render(w, r, http.StatusOK, "create", data)
}

// represents the form data and validation errors for the form fields
type snippetCreateForm struct {
	validator.Validator `schema:"-"` // ignore this field during FormData population

	Title   string `schema:"title"`
	Content string `schema:"content"`
	Expires int    `schema:"expires,default:365"`
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	var (
		form snippetCreateForm
		err  error
	)

	// // parse form data received from the client.
	// // the parsed data will be put into the r.PostForm and r.Form struct
	// if err = r.ParseForm(); err != nil {
	// 	app.clientError(w, http.StatusBadRequest)
	// 	return
	// }

	// form.Title = r.PostForm.Get("title")
	// form.Content = r.PostForm.Get("content")
	//
	// // parse "expires" into an interger before using
	// if form.Expires, err = strconv.Atoi(r.PostForm.Get("expires")); err != nil {
	// 	app.clientError(w, http.StatusBadRequest)
	// 	return
	// }

	if err := app.decodePostForm(r, &form); err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// validation cases:
	//   - Check that the "title" and "content" fields are not empty
	//   - Check that the "title" field is not more than 100 characters long
	//   - Check that the "expires" value exactly matches one of our permitted values (1, 7 or 365 days)

	form.CheckField(validator.NotBlank(form.Title), "Title", "This field cannot be blank")
	form.CheckField(validator.NotBlank(form.Content), "Content", "This field cannot be blank")

	form.CheckField(
		validator.MaxChars(form.Title, 100),
		"Title",
		"This field cannot be more than 100 characters long",
	)

	form.CheckField(
		validator.PermittedValues(form.Expires, []int{1, 7, 365}),
		"Expires",
		"This field must equal 1, 7 or 365",
	)

	// if there are any errors, rerender the same template,
	// passing in the FormErrors field
	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form

		app.render(w, r, http.StatusUnprocessableEntity, "create", data)

		return
	}

	id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
