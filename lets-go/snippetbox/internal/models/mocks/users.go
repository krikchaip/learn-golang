package mocks

import "krikchaip/snippetbox/internal/models"

var (
	dupeEmail = "dupe@example.com"
	mockUser  = struct {
		id              int
		email, password string
	}{1, "alice@example.com", "pa$$word"}
)

type UserModel struct{}

func NewUserModel() *UserModel {
	return &UserModel{}
}

func (m *UserModel) Insert(name, email, password string) error {
	switch email {
	case dupeEmail:
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	if email == mockUser.email && password == mockUser.password {
		return mockUser.id, nil
	}

	return 0, models.ErrInvalidCredentials
}

func (m *UserModel) Exists(id int) (bool, error) {
	switch id {
	case mockUser.id:
		return true, nil
	default:
		return false, nil
	}
}
