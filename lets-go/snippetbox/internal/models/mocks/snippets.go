package mocks

import (
	"krikchaip/snippetbox/internal/models"
	"time"
)

var mockSnippet = models.Snippet{
	ID:      1,
	Title:   "An old silent pond",
	Content: "An old silent pond...",
	Created: time.Now(),
	Expires: time.Now(),
}

type SnippetModel struct{}

func NewSnippetModel() *SnippetModel {
	return &SnippetModel{}
}

func (m *SnippetModel) Insert(title, content string, expires int) (id int, err error) {
	return mockSnippet.ID, nil
}

func (m *SnippetModel) Get(id int) (s models.Snippet, err error) {
	switch id {
	case mockSnippet.ID:
		return mockSnippet, nil
	default:
		return models.Snippet{}, models.ErrNoRecord
	}
}

func (m *SnippetModel) Latest() (snippets []models.Snippet, err error) {
	return []models.Snippet{mockSnippet}, nil
}
