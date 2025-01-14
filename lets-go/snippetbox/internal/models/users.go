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

type UserModelInterface interface {
	Insert(name, email, password string) error
	Authenticate(email, password string) (int, error)
	Exists(id int) (bool, error)
	Get(id int) (*User, error)
	PasswordUpdate(id int, currentPassword, newPassword string) error
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
	var (
		id             int
		hashedPassword []byte
	)

	row := m.db.QueryRow(`
		SELECT id, hashed_password FROM users
		WHERE email = $1
	`, email)

	err := row.Scan(&id, &hashedPassword)

	// when user is not found, returns ErrInvalidCredentials
	if errors.Is(err, sql.ErrNoRows) {
		return 0, ErrInvalidCredentials
	} else if err != nil {
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))

	// even if the user was found, but the password is incorrect,
	// we still return ErrInvalidCredentials anyways
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return 0, ErrInvalidCredentials
	} else if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *UserModel) Exists(id int) (bool, error) {
	var exists bool

	row := m.db.QueryRow(`SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)`, id)
	if err := row.Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}

func (m *UserModel) Get(id int) (*User, error) {
	var user User

	row := m.db.QueryRow(`
		SELECT id, name, email, created
		FROM users WHERE id = $1
	`, id)

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Created)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNoRecord
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *UserModel) PasswordUpdate(id int, currentPassword, newPassword string) error {
	var hashedPassword []byte

	row := m.db.QueryRow(`SELECT hashed_password FROM users WHERE id = $1`, id)
	err := row.Scan(&hashedPassword)

	if errors.Is(err, sql.ErrNoRows) {
		return ErrNoRecord
	} else if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(currentPassword))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return ErrInvalidCredentials
	} else if err != nil {
		return err
	}

	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), 12)
	if err != nil {
		return err
	}

	_, err = m.db.Exec(
		`UPDATE users SET hashed_password = $1 WHERE id = $2`,
		string(newHashedPassword),
		id,
	)

	return err
}
