package main

import (
	"html/template"
	"krikchaip/snippetbox/internal/models"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

type templateData struct {
	CurrentYear int
	Snippet     models.Snippet
	Snippets    []models.Snippet
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := make(map[string]*template.Template)

	pages, err := filepath.Glob("ui/html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		// the base template must be the first file!
		t, err := template.ParseFiles("ui/html/base.tmpl.html")
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

		// for example, "home.tmpl.html" -> "home"
		name := strings.Split(filepath.Base(page), ".")[0]

		// assign a template cache to the corresponding page name
		cache[name] = t
	}

	return cache, nil
}

func (app *application) newTemplateData(_ *http.Request) templateData {
	return templateData{
		CurrentYear: time.Now().Year(),
	}
}
