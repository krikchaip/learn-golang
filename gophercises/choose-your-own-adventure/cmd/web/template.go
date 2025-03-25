package main

import (
	"html/template"
	"io/fs"
)

func loadTemplate(filesystem fs.FS) *template.Template {
	return template.Must(template.ParseFS(filesystem, "*.html"))
}
