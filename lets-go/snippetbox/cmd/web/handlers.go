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

func healthz(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("OK"))
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

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	app.render(w, r, http.StatusOK, "about", data)
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

func (app *application) accountView(w http.ResponseWriter, r *http.Request) {
	id := app.sessionManager.GetInt(r.Context(), "authenticatedUserID")

	user, err := app.users.Get(id)

	if errors.Is(err, models.ErrNoRecord) {
		http.Redirect(w, r, "/user/login", http.StatusSeeOther)
	}

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := app.newTemplateData(r)
	data.User = *user

	app.render(w, r, http.StatusOK, "account", data)
}

func (app *application) userLogoutPost(w http.ResponseWriter, r *http.Request) {
	// use the RenewToken() method on the current session
	// to change the session ID again before logging out
	if err := app.sessionManager.RenewToken(r.Context()); err != nil {
		app.serverError(w, r, err)
		return
	}
	// remove the authenticatedUserID from the session data
	// so that the user is 'logged out'
	app.sessionManager.Remove(r.Context(), "authenticatedUserID")

	app.sessionManager.Put(r.Context(), "flash", "You've been logged out successfully!")

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// represents the form data and validation errors for the form fields
type snippetCreateForm struct {
	validator.Validator `schema:"-"` // ignore this field during FormData population

	Title   string `schema:"title"`
	Content string `schema:"content"`
	Expires int    `schema:"expires,default:365"`
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	// initialize template data so that it won't complain of missing values
	data.Form = snippetCreateForm{
		Expires: 365,
	}

	app.render(w, r, http.StatusOK, "create", data)
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

	// set 'flash' message of the current request context
	app.sessionManager.Put(r.Context(), "flash", "Snippet successfully created!")

	// carries the same 'r' instance during redirection
	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}

type userSignupForm struct {
	validator.Validator `schema:"-"` // ignore this field during FormData population

	Name     string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

func (app *application) userSignup(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = userSignupForm{}

	app.render(w, r, http.StatusOK, "signup", data)
}

func (app *application) userSignupPost(w http.ResponseWriter, r *http.Request) {
	var form userSignupForm

	if err := app.decodePostForm(r, &form); err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// validation cases:
	//   - Check that the provided name, email address and password are not blank.
	//   - Sanity check the format of the email address.
	//   - Ensure that the password is at least 8 characters long.
	//   - Make sure that the email address isn’t already in use.

	form.CheckField(validator.NotBlank(form.Name), "Name", "This field cannot be blank")
	form.CheckField(validator.NotBlank(form.Email), "Email", "This field cannot be blank")
	form.CheckField(validator.NotBlank(form.Password), "Password", "This field cannot be blank")

	form.CheckField(
		validator.Matches(form.Email, validator.EmailRegex),
		"Email",
		"This field must be a valid email address",
	)

	form.CheckField(
		validator.MinChars(form.Password, 8),
		"Password",
		"This field must be at least 8 characters long",
	)

	data := app.newTemplateData(r)

	if !form.Valid() {
		data.Form = form
		app.render(w, r, http.StatusUnprocessableEntity, "signup", data)

		return
	}

	err := app.users.Insert(form.Name, form.Email, form.Password)

	// validate email uniqueness and render the same form with 'email' error
	if errors.Is(err, models.ErrDuplicateEmail) {
		form.AddFieldError("Email", "Email address is already in use")

		data.Form = form
		app.render(w, r, http.StatusUnprocessableEntity, "signup", data)

		return
	}

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.sessionManager.Put(r.Context(), "flash", "Your signup was successful. Please log in.")

	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

type userLoginForm struct {
	validator.Validator `schema:"-"`

	Email    string `schema:"email"`
	Password string `schema:"password"`
}

func (app *application) userLogin(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = userLoginForm{}

	app.render(w, r, http.StatusOK, "login", data)
}

func (app *application) userLoginPost(w http.ResponseWriter, r *http.Request) {
	var form userLoginForm

	if err := app.decodePostForm(r, &form); err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// validation cases:
	//   - Check that the provided email address and password are not blank.
	//   - Sanity check the format of the email address.

	form.CheckField(validator.NotBlank(form.Email), "Email", "This field cannot be blank")
	form.CheckField(validator.NotBlank(form.Password), "Password", "This field cannot be blank")

	form.CheckField(
		validator.Matches(form.Email, validator.EmailRegex),
		"Email",
		"This field must be a valid email address",
	)

	data := app.newTemplateData(r)

	if !form.Valid() {
		data.Form = form
		app.render(w, r, http.StatusUnprocessableEntity, "login", data)

		return
	}

	id, err := app.users.Authenticate(form.Email, form.Password)

	if errors.Is(err, models.ErrInvalidCredentials) {
		form.AddNonFieldError("Email or password is incorrect")

		data.Form = form
		app.render(w, r, http.StatusUnprocessableEntity, "login", data)

		return
	}

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// NOTE: it is important to renew the session id when the authentication state
	//  or privilege levels changes for the user (e.g. login and logout operations)
	if err := app.sessionManager.RenewToken(r.Context()); err != nil {
		app.serverError(w, r, err)
		return
	}

	redirectPathAfterLogin := "/snippet/create"

	app.sessionManager.Put(r.Context(), "authenticatedUserID", id)

	// use the PopString() method to retrieve and remove a value from the session data in one step
	if path := app.sessionManager.PopString(r.Context(), "redirectPathAfterLogin"); path != "" {
		redirectPathAfterLogin = path
	}

	// Redirect the user to the create snippet page
	http.Redirect(w, r, redirectPathAfterLogin, http.StatusSeeOther)
}

type accountPasswordUpdateForm struct {
	validator.Validator `schema:"-"`

	CurrentPassword         string `schema:"currentPassword"`
	NewPassword             string `schema:"newPassword"`
	NewPasswordConfirmation string `schema:"newPasswordConfirmation"`
}

func (app *application) accountPasswordUpdate(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = accountPasswordUpdateForm{}

	app.render(w, r, http.StatusOK, "password", data)
}

func (app *application) accountPasswordUpdatePost(w http.ResponseWriter, r *http.Request) {
	var form accountPasswordUpdateForm

	if err := app.decodePostForm(r, &form); err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// validation cases:
	//   - All three fields are required
	//   - The newPassword value must be at least 8 characters long
	//   - The newPassword and newPasswordConfirmation values must match

	form.CheckField(
		validator.NotBlank(form.CurrentPassword),
		"CurrentPassword",
		"This field cannot be blank",
	)
	form.CheckField(
		validator.NotBlank(form.NewPassword),
		"NewPassword",
		"This field cannot be blank",
	)
	form.CheckField(
		validator.NotBlank(form.NewPasswordConfirmation),
		"NewPasswordConfirmation",
		"This field cannot be blank",
	)

	form.CheckField(
		validator.MinChars(form.NewPassword, 8),
		"NewPassword",
		"This field must be at least 8 characters long",
	)

	form.CheckField(
		form.NewPassword == form.NewPasswordConfirmation,
		"NewPasswordConfirmation",
		"Passwords do not match",
	)

	data := app.newTemplateData(r)

	if !form.Valid() {
		data.Form = form
		app.render(w, r, http.StatusUnprocessableEntity, "password", data)

		return
	}

	id := app.sessionManager.GetInt(r.Context(), "authenticatedUserID")
	err := app.users.PasswordUpdate(id, form.CurrentPassword, form.NewPassword)

	if errors.Is(err, models.ErrInvalidCredentials) {
		form.AddFieldError("CurrentPassword", "Current password is incorrect")

		data.Form = form
		app.render(w, r, http.StatusUnprocessableEntity, "password", data)

		return
	}

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.sessionManager.Put(r.Context(), "flash", "Your password has been updated!")
	http.Redirect(w, r, "/account/view", http.StatusSeeOther)
}
