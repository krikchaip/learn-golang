package models

import (
	"database/sql"
	"errors"
	"time"
)

// Define a Snippet type to hold the data for an individual snippet
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// Define a SnippetModel type which wraps a sql.DB connection pool.
type SnippetModel struct {
	db *sql.DB
}

func NewSnippetModel(db *sql.DB) *SnippetModel {
	return &SnippetModel{db}
}

// This will insert a new snippet into the database.
func (m *SnippetModel) Insert(title, content string, expires int) (id int, err error) {
	// postgresql does not support Exec().LastInsertId() so we have to find another way
	// ref: https://stackoverflow.com/questions/33382981/go-how-to-get-last-insert-id-on-postgresql-with-namedexec
	//      https://stackoverflow.com/questions/71378392/go-postgres-prepared-statement-with-interval-parameter-not-working
	row := m.db.QueryRow(`
		INSERT INTO snippets (title, content, expires)
		VALUES ($1, $2, NOW() + INTERVAL '1 day' * $3)
		RETURNING id
	`, title, content, expires)

	err = row.Err()
	if err != nil {
		return 0, err
	}

	// load the result into the 'id' variable
	if err = row.Scan(&id); err != nil {
		return 0, err
	}

	return
}

// This will return a specific snippet based on its id.
func (m *SnippetModel) Get(id int) (s Snippet, err error) {
	// errors are deferred until Row's Scan method is called
	row := m.db.QueryRow(`
		SELECT id, title, content, created, expires
		FROM snippets WHERE expires > NOW() AND id = $1
	`, id)

	err = row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

	if errors.Is(err, sql.ErrNoRows) {
		return Snippet{}, ErrNoRecord
	}

	if err != nil {
		return Snippet{}, err
	}

	return s, nil
}

// This will return the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]Snippet, error) {
	return nil, nil
}
