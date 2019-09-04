package app

import (
	"errors"

	"github.com/deni1688/motusauth/models"
	"github.com/stretchr/testify/mock"
)

type MockApp struct {
	mock.Mock
}

func (a *MockApp) AuthenticateUser(u models.User) (string, error) {
	if u.Password == "" || u.Email == "" {
		return "", errors.New("Email and Password are required")
	}

	testPassword := hashAndSalt("testing123")

	if err := comparePasswords(testPassword, []byte(u.Password)); err != nil {
		return "", err
	}

	return "mockToken123", nil
}

// RegisterUser ...
func (a *MockApp) RegisterUser(u *models.User) error {
	return ValidateUser(u)
}
