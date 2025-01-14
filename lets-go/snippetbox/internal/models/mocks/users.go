package mocks

import (
	"krikchaip/snippetbox/internal/models"
	"time"
)

var (
	DupeEmail = "dupe@example.com"
	MockUser  = struct {
		Id                    int
		Name, Email, Password string
	}{1, "alice", "alice@example.com", "pa$$word"}
)

type UserModel struct{}

func NewUserModel() *UserModel {
	return &UserModel{}
}

func (m *UserModel) Insert(name, email, password string) error {
	switch email {
	case DupeEmail:
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	if email == MockUser.Email && password == MockUser.Password {
		return MockUser.Id, nil
	}

	return 0, models.ErrInvalidCredentials
}

func (m *UserModel) Exists(id int) (bool, error) {
	switch id {
	case MockUser.Id:
		return true, nil
	default:
		return false, nil
	}
}

func (m *UserModel) Get(id int) (*models.User, error) {
	user := models.User{
		ID:      mockSnippet.ID,
		Name:    MockUser.Name,
		Email:   MockUser.Email,
		Created: time.Now(),
	}

	switch id {
	case MockUser.Id:
		return &user, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *UserModel) PasswordUpdate(id int, currentPassword, newPassword string) error {
	if id != MockUser.Id {
		return models.ErrNoRecord
	}

	if currentPassword != MockUser.Password {
		return models.ErrInvalidCredentials
	}

	return nil
}
