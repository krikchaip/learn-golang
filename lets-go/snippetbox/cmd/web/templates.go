package main

import (
	"html/template"
	"krikchaip/snippetbox/internal/models"
	"path/filepath"
	"strings"
)

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := make(map[string]*template.Template)

	pages, err := filepath.Glob("ui/html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		files := []string{
			"ui/html/base.tmpl.html", // the base template must be the first file!
			"ui/html/partials/nav.tmpl.html",
			page,
		}

		t, err := template.ParseFiles(files...)
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
