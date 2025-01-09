package models

import "errors"

var (
	ErrNoRecord = errors.New("models: no matching record found")

	// use when a user tries to login with an incorrect email or password
	ErrInvalidCredentials = errors.New("models: invalid credentials")

	// use when a user tries to signup with an email that's already in use
	ErrDuplicateEmail = errors.New("models: duplicate email")
)
