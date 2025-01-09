package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

type UserModel struct {
	db *sql.DB
}

func NewUserModel(db *sql.DB) *UserModel {
	return &UserModel{db}
}

func (m *UserModel) Insert(name, email, password string) error {
	// a cost of '12' means that 4096 (2^12) bcrypt iterations
	// will be used to generate the password hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	_, err = m.db.Exec(`
		INSERT INTO users (name, email, hashed_password)
		VALUES ($1, $2, $3)
	`, name, email, string(hashedPassword))

	// the error returned above will have type *pgconn.PgError
	// ref: https://github.com/jackc/pgx/wiki/Error-Handling
	var pgError *pgconn.PgError

	// we need to check whether or not the error relates to
	// violating unique constraint for the 'email' key
	// ref: https://pkg.go.dev/github.com/jackc/pgconn#PgError
	if errors.As(err, &pgError) && pgError.Code == "23505" &&
		pgError.ConstraintName == "users_email_key" {
		return ErrDuplicateEmail
	}

	if err != nil {
		return err
	}

	return nil
}

// returns user's ID if success
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

func (m *UserModel) Exists(id int) (bool, error) {
	return false, nil
}
