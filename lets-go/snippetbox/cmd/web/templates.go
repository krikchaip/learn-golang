package main

import (
	"html/template"
	"krikchaip/snippetbox/internal/models"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

// functions that will be used in the template files
var functions = template.FuncMap{
	"humanDate": func(t time.Time) string {
		return t.Format("2 Jan 2006 at 15:04")
	},
}

type templateData struct {
	CurrentYear int
	Snippet     models.Snippet
	Snippets    []models.Snippet
}

func (app *application) newTemplateData(_ *http.Request) templateData {
	return templateData{
		CurrentYear: time.Now().Year(),
	}
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := make(map[string]*template.Template)

	pages, err := filepath.Glob("ui/html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		// for example, "home.tmpl.html" -> "home"
		name := strings.Split(filepath.Base(page), ".")[0]

		// template functions must be registered before calling template.Parse*()
		t := template.New(name).Funcs(functions)

		// the base template must be the first file!
		t, err := t.ParseFiles("ui/html/base.tmpl.html")
		if err != nil {
			return nil, err
		}

		// Call ParseGlob() *on this template set* to add any partials.
		t, err = t.ParseGlob("ui/html/partials/*.tmpl.html")
		if err != nil {
			return nil, err
		}

		// parse the page template the last. this is to make sure that
		// all necessary templates are compiled
		t, err = t.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// assign a template cache to the corresponding page name
		cache[name] = t
	}

	return cache, nil
}
