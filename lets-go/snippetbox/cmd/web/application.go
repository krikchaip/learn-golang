package main

import (
	"html/template"
	"krikchaip/snippetbox/internal/models"
	"log/slog"
)

// the application's dependencies (for dependency injection)
type application struct {
	logger        *slog.Logger
	snippets      *models.SnippetModel
	templateCache map[string]*template.Template
}
