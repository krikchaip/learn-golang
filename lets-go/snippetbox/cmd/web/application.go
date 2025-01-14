package main

import (
	"html/template"
	"krikchaip/snippetbox/internal/models"
	"log/slog"

	"github.com/alexedwards/scs/v2"
	"github.com/gorilla/schema"
)

// the application's dependencies (for dependency injection)
type application struct {
	debug bool // whether the debug mode is enabled

	logger         *slog.Logger
	templateCache  map[string]*template.Template
	decoder        *schema.Decoder
	sessionManager *scs.SessionManager

	snippets models.SnippetModelInterface
	users    models.UserModelInterface
}
