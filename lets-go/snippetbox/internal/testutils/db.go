package testutils

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var (
	dsn         = "postgresql://postgres:postgres@localhost:5432/snippetbox?search_path=test"
	setupSQL    = "../models/testdata/setup.sql"
	teardownSQL = "../models/testdata/teardown.sql"
)

func NewTestDB(t *testing.T) *sql.DB {
	// ref: https://github.com/jackc/pgx/wiki/Getting-started-with-pgx-through-database-sql
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		t.Fatal(err)
	}

	if err := executeSQLScript(db, setupSQL); err != nil {
		t.Fatal(err)
	}

	// the function passed to the Cleanup() method will automatically be called
	// when the current test (or sub-test) has finished.
	t.Cleanup(func() {
		if err := executeSQLScript(db, teardownSQL); err != nil {
			t.Fatal(err)
		}
	})

	return db
}

func executeSQLScript(db *sql.DB, path string) error {
	// read the setup SQL script from the file and execute the statements,
	// closing the connection pool in the event of an error.
	script, err := os.ReadFile(path)
	if err != nil {
		db.Close()
		return err
	}

	_, err = db.Exec(string(script))
	if err != nil {
		db.Close()
		return err
	}

	return nil
}
