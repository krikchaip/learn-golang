package ui

import "embed"

// embedding files/folders into package variables
// NOTE:
//   - the path must be relative to this .go file
//   - can only be use on global variables at package level
//   - paths can not contain . or .. elements
//   - the embed FS is _always_ rooted in the directory containing the .go file

//go:embed "static"
var StaticFiles embed.FS

//go:embed "html/*"
var TemplateFiles embed.FS
